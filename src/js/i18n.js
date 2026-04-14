/**
 * UI strings from WASM (go-i18n + embedded YAML). Fallback: key id.
 */

const STORAGE_KEY = 'lux_locale';

const localeListeners = new Set();

let dict = Object.create(null);

function browserDefaultTag() {
    if (typeof navigator === 'undefined') {
        return 'en';
    }
    const list = navigator.languages && navigator.languages.length
        ? navigator.languages
        : [navigator.language || 'en'];
    for (let i = 0; i < list.length; i++) {
        const x = String(list[i] || '').toLowerCase();
        if (x.startsWith('ru')) {
            return 'ru';
        }
    }
    return 'en';
}

export function resolveLocaleTag() {
    try {
        const s = String(localStorage.getItem(STORAGE_KEY) || '').toLowerCase();
        if (s === 'ru' || s === 'en') {
            return s;
        }
    } catch (_) {
        /* ignore */
    }
    return browserDefaultTag();
}

/** "ru" | "en" — synced when applyLocale runs; initial value follows storage / browser. */
let currentTag = resolveLocaleTag();

/** Suffixes for incoming damage source column; rebuilt when locale or WASM dict changes. */
let incomingSourceSuffixStripList = null;

function buildIncomingSourceSuffixStripList() {
    const fn = globalThis.getUITranslationsJSON;
    const set = new Set([' (receives)', ' (получает)']);
    if (typeof fn === 'function') {
        ['en', 'ru'].forEach((lang) => {
            try {
                const o = JSON.parse(fn(lang) || '{}');
                const s = o.df_source_incoming_suffix;
                if (typeof s === 'string' && s !== '') {
                    set.add(s);
                }
            } catch (_) {
                /* ignore */
            }
        });
    }
    return Array.from(set).sort((a, b) => b.length - a.length);
}

/** Longest-first suffixes used to normalize damage source labels across locale switches. */
export function incomingSourceSuffixVariants() {
    if (!incomingSourceSuffixStripList) {
        incomingSourceSuffixStripList = buildIncomingSourceSuffixStripList();
    }
    return incomingSourceSuffixStripList;
}

export function setStoredLocaleTag(tag) {
    const norm = tag === 'ru' ? 'ru' : 'en';
    try {
        localStorage.setItem(STORAGE_KEY, norm);
    } catch (_) {
        /* ignore */
    }
    return norm;
}

export function getLocaleTag() {
    return currentTag;
}

/** BCP 47 for Intl.* */
export function getBcp47Locale() {
    return currentTag === 'ru' ? 'ru-RU' : 'en-US';
}

export function subscribeLocale(cb) {
    localeListeners.add(cb);
    return () => localeListeners.delete(cb);
}

function notifyLocale() {
    localeListeners.forEach((cb) => {
        try {
            cb(currentTag);
        } catch (_) {
            /* ignore */
        }
    });
}

function fetchDictFromWasm(lang) {
    const fn = globalThis.getUITranslationsJSON;
    if (typeof fn !== 'function') {
        return null;
    }
    try {
        const raw = fn(lang);
        const o = JSON.parse(raw || '{}');
        if (!o || typeof o !== 'object') {
            return null;
        }
        return o;
    } catch (_) {
        return null;
    }
}

export function t(id) {
    const v = dict[id];
    return v != null && v !== '' ? v : id;
}

/** Replace `{name}` placeholders in a translated string. */
export function tf(id, vars) {
    let s = t(id);
    if (!vars) {
        return s;
    }
    Object.keys(vars).forEach((k) => {
        s = s.split(`{${k}}`).join(String(vars[k]));
    });
    return s;
}

function applyDomFromDict() {
    document.documentElement.lang = currentTag === 'ru' ? 'ru' : 'en';

    document.querySelectorAll('[data-i18n]').forEach((el) => {
        const id = el.getAttribute('data-i18n');
        if (!id) {
            return;
        }
        const val = t(id);
        if (el.hasAttribute('data-i18n-html')) {
            el.innerHTML = val;
        } else {
            el.textContent = val;
        }
    });

    document.querySelectorAll('[data-i18n-placeholder]').forEach((el) => {
        const id = el.getAttribute('data-i18n-placeholder');
        if (!id || !('placeholder' in el)) {
            return;
        }
        el.placeholder = t(id);
    });

    document.querySelectorAll('[data-i18n-title]').forEach((el) => {
        const id = el.getAttribute('data-i18n-title');
        if (!id) {
            return;
        }
        el.title = t(id);
    });

    document.querySelectorAll('[data-i18n-aria]').forEach((el) => {
        const id = el.getAttribute('data-i18n-aria');
        if (!id) {
            return;
        }
        el.setAttribute('aria-label', t(id));
    });

    document.title = t('page_title');
}

/**
 * Load dictionary for lang, set globals, apply [data-i18n*].
 * @param {string} lang "ru" | "en"
 * @param {{ skipDom?: boolean }} opts
 */
export function applyLocale(lang, opts = {}) {
    const tag = lang === 'ru' ? 'ru' : 'en';
    const syncGo = globalThis.setUILanguage;
    if (typeof syncGo === 'function') {
        try {
            syncGo(tag);
        } catch (_) {
            /* ignore */
        }
    }
    incomingSourceSuffixStripList = null;
    const next = fetchDictFromWasm(tag) || {};
    dict = next;
    currentTag = tag;
    if (!opts.skipDom) {
        applyDomFromDict();
    }
    notifyLocale();
    return tag;
}

export function initLocaleSwitcher() {
    const wrap = document.getElementById('lang_switch');
    if (!wrap || wrap.dataset.luxI18nInit === '1') {
        return;
    }
    wrap.dataset.luxI18nInit = '1';
    const syncActive = () => {
        wrap.querySelectorAll('button[data-locale]').forEach((b) => {
            const on = b.getAttribute('data-locale') === currentTag;
            b.classList.toggle('is-active', on);
            b.setAttribute('aria-pressed', on ? 'true' : 'false');
        });
    };
    wrap.querySelectorAll('button[data-locale]').forEach((btn) => {
        btn.addEventListener('click', () => {
            const next = btn.getAttribute('data-locale');
            if (next !== 'ru' && next !== 'en') {
                return;
            }
            setStoredLocaleTag(next);
            applyLocale(next);
            syncActive();
        });
    });
    syncActive();
}

export function bootstrapI18nAfterWasm() {
    const tag = resolveLocaleTag();
    applyLocale(tag);
    initLocaleSwitcher();
}
