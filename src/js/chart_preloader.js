/**
 * Прелоадеры поверх тяжёлых графиков: даём браузеру отрисовать спиннер до WASM/рендера.
 */

const LOADER_CLASS = 'chart-preloader';
const SPINNER_CLASS = 'chart-preloader-spinner';

/** Два rAF — типичный способ дождаться следующего кадра после изменений DOM. */
export function deferAfterPaint(fn) {
    if (typeof requestAnimationFrame !== 'function') {
        setTimeout(fn, 0);
        return;
    }
    requestAnimationFrame(() => {
        requestAnimationFrame(fn);
    });
}

export function showChartPreloader(host, options = {}) {
    if (!host) {
        return null;
    }
    const label = options.label || 'Загрузка…';
    host.querySelectorAll(`:scope > .${LOADER_CLASS}`).forEach((el) => el.remove());
    const wrap = document.createElement('div');
    wrap.className = LOADER_CLASS;
    wrap.setAttribute('role', 'status');
    wrap.setAttribute('aria-live', 'polite');
    wrap.setAttribute('aria-busy', 'true');
    wrap.setAttribute('aria-label', label);
    const sp = document.createElement('div');
    sp.className = SPINNER_CLASS;
    wrap.appendChild(sp);
    host.appendChild(wrap);
    return wrap;
}

export function hideChartPreloader(host) {
    if (!host) {
        return;
    }
    host.querySelectorAll(`:scope > .${LOADER_CLASS}`).forEach((el) => el.remove());
}
