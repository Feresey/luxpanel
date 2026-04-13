/**
 * GA4 (gtag) — только если в шаблоне подключён googletagmanager/gtag.
 * События с префиксом lux_* — удобно фильтровать в отчётах; при необходимости
 * зарегистрируйте параметры как custom dimensions / metrics в Admin → Data display.
 *
 * Тайминги (мс, округлённые):
 * - lux_logs_parsed: log_parse_ms, match_list_ms, parse_total_ms
 * - lux_charts_timing: get_charts_wasm_ms, charts_json_ms, get_matrices_wasm_ms, matrices_json_ms,
 *   charts_data_prep_ms (сумма wasm+json), pie_render_ms, matrix_dom_ms, charts_render_ms (пироги+таблицы),
 *   js_heap_used_before_mb / js_heap_used_after_mb (если performance.memory — обычно Chrome; иначе 0)
 * - lux_logs_parsed: опционально js_heap_used_after_mb после parseFiles
 * - lux_timeline_timing, lux_battle_charts_timing — аналогично по блокам
 */

export function gaEvent(name, params = {}) {
    if (typeof window === 'undefined') {
        return;
    }
    const g = window.gtag;
    if (typeof g !== 'function') {
        return;
    }
    try {
        g('event', name, params);
    } catch (_) {
        /* ignore */
    }
}

/** Грубые корзины размера (без точных байт — меньше шума и чувствительных данных). */
export function byteSizeBucket(byteLen) {
    const n = Number(byteLen) || 0;
    if (n < 50_000) {
        return 'xs';
    }
    if (n < 200_000) {
        return 's';
    }
    if (n < 1_000_000) {
        return 'm';
    }
    if (n < 5_000_000) {
        return 'l';
    }
    return 'xl';
}
