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
                    'rgba(255, 159, 64, 0.2)'
                ],
                borderColor: [
                    'rgba(0, 0, 0, 0.5)',
                    'rgba(0, 0, 0, 0.5)',
                    'rgba(0, 0, 0, 0.5)',
                    'rgba(0, 0, 0, 0.5)',
                    'rgba(0, 0, 0, 0.5)',
                    'rgba(0, 0, 0, 0.5)',
                    'rgba(0, 0, 0, 0.5)'
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
                text: 'Category Distribution',
                display: true,
                lineHeight: 2.5,
                fontSize: 25,
                fontColor: 'white',
            },
            legend: {
                position: 'right',
                labels: {
                    boxWidth: 50,
                    fontSize: 18,
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