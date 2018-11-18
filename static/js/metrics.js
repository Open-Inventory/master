$(document).ready(function() {
    var ctx = document.getElementById("categoryCountCharts").getContext('2d');
    var categoryCountCharts = new Chart(ctx, {
        type: 'polarArea',
        data: {
            labels: ["Keyboards/Mice", "Motherboards", "Processors", "Video Cards", "RAM", "Monitors", "Hard Drives/SSDs"],
            datasets: [{
                // label: '# of Products',
                data: [13, 19, 8, 14, 17, 11, 20],
                backgroundColor: [
                    'rgba(255, 125, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(255, 206, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(255, 159, 64, 0.2)',
                    'rgba(50, 50, 50, 0.2)'
                ],
                borderColor: [
                    'rgba(255,99,132,1)',
                    'rgba(54, 162, 235, 1)',
                    'rgba(255, 206, 86, 1)',
                    'rgba(75, 192, 192, 1)',
                    'rgba(153, 102, 255, 1)',
                    'rgba(255, 159, 64, 1)',
                    'rgba(50, 50, 50, 1)'
                ],
                borderWidth: 1,
                hoverBorderWidth: 2,
            }]
        },
        options: {
            animation: {
                animateScale: false,
            },
            tooltips: {
                backgroundColor: 'orange',
                bodyFontSize: 14,
            },
            title: {
                text: 'Product Distribution',
                display: true,
                lineHeight: 2.5,
                fontSize: 25,
                fontColor: 'white',
            },
            legend: {
                position: 'right',
                labels: {
                    boxWidth: 50,
                    fontSize: 15,
                    fontColor: 'white',
                    padding: 15,
                }
            },
            scale: {
                ticks: {
                    min: 0,
                    maxTicksLimit: 5,
                    fontSize: 15,
                    fontColor: 'white',
                    showLabelBackdrop: false
                },
                gridLines: {
                    lineWidth: 0.5
                }
            },
            
        }
    });
    var ctx1 = document.getElementById("categoryValueChart").getContext('2d');
    var categoryValueChart = new Chart(ctx1, {
        type: 'polarArea',
        data: {
            labels: ["Keyboards/Mice", "Motherboards", "Processors", "Video Cards", "RAM", "Monitors", "Hard Drives/SSDs"],
            datasets: [{
                // label: '# of Products',
                data: [1510, 2223, 3521, 3123, 2145, 1206, 2541],
                backgroundColor: [
                    'rgba(255, 125, 132, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(255, 206, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(255, 159, 64, 0.2)',
                    'rgba(50, 50, 50, 0.2)'
                ],
                borderColor: [
                    'rgba(255,99,132,1)',
                    'rgba(54, 162, 235, 1)',
                    'rgba(255, 206, 86, 1)',
                    'rgba(75, 192, 192, 1)',
                    'rgba(153, 102, 255, 1)',
                    'rgba(255, 159, 64, 1)',
                    'rgba(50, 50, 50, 1)'
                ],
                borderWidth: 1,
                hoverBorderWidth: 2,
            }]
        },
        options: {
            animation: {
                animateScale: false,
            },
            tooltips: {
                backgroundColor: 'orange',
                bodyFontSize: 14,
                callbacks: {
                    label: function(tooltipItems, data) {

                        return  data.labels[tooltipItems.index] + ": $" + tooltipItems.yLabel.toString();
                    }
                }
                
            },
            title: {
                text: 'Category Value Distribution',
                display: true,
                lineHeight: 2.5,
                fontSize: 25,
                fontColor: 'white',
            },
            legend: {
                position: 'right',
                labels: {
                    boxWidth: 50,
                    fontSize: 15,
                    fontColor: 'white',
                    padding: 15,
                }
            },
            scale: {
                ticks: {
                    min: 0,
                    maxTicksLimit: 5,
                    fontSize: 15,
                    fontColor: 'white',
                    showLabelBackdrop: false,
                    callback: function(value, index, values) {
                        return '$' + value;
                    }
                },
                gridLines: {
                    lineWidth: 0.5
                }
            },
            
        }
    });
});