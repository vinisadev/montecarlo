// Initialize the chart
const ctx = document.getElementById('myChart').getContext('2d');
let myChart = new Chart(ctx, {
    type: 'line',
    data: {
        labels: [],
        datasets: [{
            label: 'Percentage Over Days',
            data: [],
            borderColor: 'rgba(75, 192, 192, 1)',
            borderWidth: 1,
            fill: false
        }]
    },
    options: {
        scales: {
            y: {
                beginAtZero: true,
                max: 100,
                title: {
                    display: true,
                    text: 'Percentage'
                }
            },
            x: {
                title: {
                    display: true,
                    text: 'Days'
                }
            }
        }
    }
});

// Function to update the chart based on JSON input
function updateChart() {
    const jsonInput = document.getElementById('jsonInput').value;
    try {
        const data = JSON.parse(jsonInput);
        if (Array.isArray(data)) {
            const labels = data.map(item => item.days);
            const percentages = data.map(item => item.percentage);

            myChart.data.labels = labels;
            myChart.data.datasets[0].data = percentages;
            myChart.update();
        } else {
            alert('JSON data must be an array of objects with "days" and "percentage" properties.');
        }
    } catch (error) {
        alert('Invalid JSON data.');
    }
}
