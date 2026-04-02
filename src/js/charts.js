import { Chart, registerables } from "chart.js";
import ChartDataLabels from "chartjs-plugin-datalabels";

Chart.register(...registerables, ChartDataLabels);

export { CreateCharts, ApplyParsedCharts };

const pieBorder = 'rgba(24, 24, 28, 0.55)';
const pieBorderWidth = 1.5;

const pieColors = [
    'rgba(56, 189, 248, 0.92)',
    'rgba(167, 139, 250, 0.92)',
    'rgba(52, 211, 153, 0.92)',
    'rgba(251, 191, 36, 0.92)',
    'rgba(251, 113, 133, 0.92)',
    'rgba(45, 212, 191, 0.92)',
    'rgba(251, 146, 60, 0.92)',
    'rgba(232, 121, 249, 0.92)',
    'rgba(148, 163, 184, 0.92)',
    'rgba(96, 165, 250, 0.92)',
];

function chartDevicePixelRatio() {
    if (typeof window === "undefined" || !window.devicePixelRatio) {
        return 1;
    }
    return Math.min(Math.max(window.devicePixelRatio, 1), 2.5);
}

const pieOptions = {
    responsive: true,
    maintainAspectRatio: true,
    aspectRatio: 1,
    devicePixelRatio: chartDevicePixelRatio(),
    interaction: {
        mode: 'nearest',
        intersect: true,
    },
    plugins: {
        legend: {
            display: false,
        },
        datalabels: {
            display: (ctx) => {
                const v = ctx.dataset.data[ctx.dataIndex];
                const n = typeof v === "number" ? v : Number(v);
                return Number.isFinite(n) && n > 0;
            },
            formatter: (_value, ctx) => {
                const label = ctx.chart.data.labels[ctx.dataIndex];
                return label != null ? String(label) : "";
            },
            color: "#f8fafc",
            font: {
                size: 12,
                weight: "600",
                family: 'system-ui, "Segoe UI", Roboto, sans-serif',
            },
            textStrokeColor: "rgba(15, 23, 42, 0.4)",
            textStrokeWidth: 1.25,
        },
        tooltip: {
            enabled: true,
            backgroundColor: 'rgba(15, 23, 42, 0.92)',
            titleColor: '#e2e8f0',
            bodyColor: '#f1f5f9',
            borderColor: 'rgba(148, 163, 184, 0.35)',
            borderWidth: 1,
            padding: 10,
            displayColors: true,
            callbacks: {
                title() {
                    return '';
                },
                label(ctx) {
                    const name = ctx.label || '';
                    const v = ctx.raw;
                    const n = typeof v === 'number' ? v : Number(v);
                    const s = Number.isInteger(n) ? String(n) : n.toFixed(1);
                    return `${name}: ${s}`;
                },
            },
        },
    },
};

function createPieWithData(ctx, data) {
    const pie_config = {
        type: 'pie',
        data: data,
        options: pieOptions,
    };
    return new Chart(ctx, pie_config);
}

let pieChart1 = null;
let pieChart2 = null;

function getLastValue(curve) {
    if (!curve || !Array.isArray(curve.data) || curve.data.length === 0) {
        return 0;
    }
    return Number(curve.data[curve.data.length - 1]) || 0;
}

function updateChartDataset(chart, labels, values, datasetLabel) {
    if (!chart || !chart.data || !chart.data.datasets || chart.data.datasets.length === 0) {
        return;
    }
    const colors = labels.map((_, i) => pieColors[i % pieColors.length]);
    const borders = labels.map(() => pieBorder);
    chart.data.labels = labels;
    chart.data.datasets[0].label = datasetLabel;
    chart.data.datasets[0].data = values;
    chart.data.datasets[0].backgroundColor = colors;
    chart.data.datasets[0].borderColor = borders;
    chart.data.datasets[0].borderWidth = pieBorderWidth;
    chart.data.datasets[0].hoverBorderWidth = 2;
    chart.update('none');
}

function fetchChartsJSON(levelIndex, mode) {
    const g = globalThis.getChartsJSON;
    if (typeof g === 'function') {
        const raw = g(levelIndex, String(mode));
        if (raw && raw !== 'null') {
            return raw;
        }
    }
    const d = globalThis.getDamageChartsJSON;
    if (mode === 'damage' && typeof d === 'function') {
        return d(levelIndex);
    }
    return null;
}

function ApplyParsedCharts(levelIndex = 0, mode = 'damage') {
    const raw = fetchChartsJSON(levelIndex, mode);
    if (!raw || raw === 'null') {
        return;
    }

    let parsed;
    try {
        parsed = JSON.parse(raw);
    } catch (_) {
        return;
    }
    if (!Array.isArray(parsed) || parsed.length < 2) {
        return;
    }

    const team1 = Array.isArray(parsed[0]) ? parsed[0] : [];
    const team2 = Array.isArray(parsed[1]) ? parsed[1] : [];
    const t1PlayerCurves = team1.slice(1);
    const t2PlayerCurves = team2.slice(1);

    const labels1 = t1PlayerCurves.map((c) => c.name || "Unknown");
    const values1 = t1PlayerCurves.map((c) => getLastValue(c));
    const labels2 = t2PlayerCurves.map((c) => c.name || "Unknown");
    const values2 = t2PlayerCurves.map((c) => getLastValue(c));

    const label = mode === 'heal' ? 'Heal' : mode === 'kill' ? 'Kills' : 'Damage';
    updateChartDataset(pieChart1, labels1, values1, label);
    updateChartDataset(pieChart2, labels2, values2, label);
}

function makePiePlaceholder() {
    const bg = pieColors.slice(0, 3);
    return {
        labels: ['—', '—', '—'],
        datasets: [
            {
                label: 'Damage',
                backgroundColor: bg,
                borderColor: bg.map(() => pieBorder),
                borderWidth: pieBorderWidth,
                hoverOffset: 6,
                data: [1, 1, 1],
            },
        ],
    };
}

function CreateCharts() {
    pieChart1 = createPieWithData(
        document.getElementById('pieChart1').getContext('2d'),
        makePiePlaceholder()
    );
    pieChart2 = createPieWithData(
        document.getElementById('pieChart2').getContext('2d'),
        makePiePlaceholder()
    );
}
