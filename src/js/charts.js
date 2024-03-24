import { Chart, registerables } from "chart.js";

Chart.register(...registerables);

export { CreateCharts };

function CreateCharts() {
    const left_data = {
        labels: [
            'Red',
            'Blue',
            'Yellow'
        ],
        datasets: [{
            label: 'My First Dataset',
            data: [300, 50, 100],
            backgroundColor: [
                'rgb(255, 99, 132)',
                'rgb(54, 162, 235)',
                'rgb(255, 205, 86)'
            ],
            hoverOffset: 4
        }]
    };
    const left_config = {
        type: 'pie',
        data: left_data,
    };

    var pieChart1 = new Chart(document.getElementById('pieChart1').getContext('2d'), left_config);
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