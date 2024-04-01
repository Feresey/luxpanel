import { Chart, registerables } from "chart.js";
import ChartDataLabels from 'chartjs-plugin-datalabels';

Chart.register(...registerables);

Chart.defaults.set('plugins.datalabels', {
    color: '#FE777B'
});


export { CreateCharts };

function createPieWithData(ctx, data) {
    const pie_config = {
        type: 'pie',
        data: data,
        plugins: [ChartDataLabels],
        options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
                datalabels: {
                    formatter: (value, ctx) => {
                        const label = ctx.chart.data.labels[ctx.dataIndex];
                        return `${label}: ${value}`;
                    },
                    color: 'white',
                    font: {
                        size: 12,
                        weight: 'bold',
                        family: 'Arial',
                        style: 'normal',
                        lineHeight: 1,
                        align: 'center',
                        verticalAlign: 'bottom',
                    },
                },
            },
        }
    }
    return new Chart(ctx, pie_config)
}

function CreateCharts() {
    const left_data = {
        labels: [
            'Red',
            'Blue',
            'Yellow'
        ],
        datasets: [
            {
                label: "Damage",
                backgroundColor: [
                    'rgb(255, 99, 132)',
                    'rgb(54, 162, 235)',
                    'rgb(255, 205, 86)'
                ],
                data: [300, 50, 100],
                hoverOffset: 4,
            },
            {
                label: "Heal",
                backgroundColor: [
                    'rgb(255, 99, 132)',
                    'rgb(54, 162, 235)',
                    'rgb(255, 205, 86)'
                ],
                data: [200, 150, 50],
                hoverOffset: 4,
                showLine: true,
            },
            {
                label: "Kill",
                backgroundColor: [
                    'rgb(255, 99, 132)',
                    'rgb(54, 162, 235)',
                    'rgb(255, 205, 86)'
                ],
                data: [100, 200, 150],
                hoverOffset: 4,
            },
        ],
    };

    var pieChart1 = createPieWithData(document.getElementById('pieChart1').getContext('2d'), left_data)
    var pieChart2 = createPieWithData(document.getElementById('pieChart2').getContext('2d'), left_data)

    // var pieChart1 = new Chart(document.getElementById('pieChart1').getContext('2d'), left_config);
    //  {
    //     type: 'pie',
    //     data: {
    //         labels: ['Label 1', 'Label 2', 'Label 3'],
    //         datasets: [{
    //             label: 'Dataset 1',
    //             data: [10, 20, 30],
    //             backgroundColor: ['#3e95cd', '#8e5ea2', '#3cba9f'],
    //         }]
    //     },
    //     options: {
    //         responsive: true,
    //         maintainAspectRatio: false,
    //     }
    // });

    // var pieChart2 = new Chart(document.getElementById('pieChart2').getContext('2d'), {
    //     type: 'pie',
    //     data: {
    //         labels: ['Label 1', 'Label 2', 'Label 3'],
    //         datasets: [{
    //             label: 'Dataset 2',
    //             data: [40, 50, 60],
    //             backgroundColor: ['#3e95cd', '#8e5ea2', '#3cba9f'],
    //         }]
    //     },
    //     options: {
    //         responsive: true,
    //         maintainAspectRatio: false,
    //     }
    // });
}