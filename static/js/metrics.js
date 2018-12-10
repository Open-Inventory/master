$(document).ready(function() {
    $.getJSON("/api/items", function(data) {
        var items = [];
        var labels= [];
        $.each(data, function (i) {
            items.push(data[i].Quantity);
            console.log (items);
            labels.push(data[i].Name);
        })
        charts('categoryCountCharts', labels, items);
      });

    $.getJSON("/api/orders/quantity", function(data) {
    var id = [];
    var quantity= [];
    $.each(data, function (i) {
        id.push(data[i].Id);
        quantity.push(data[i].Quantity);
        console.log (quantity);

    })
    charts('categoryValueChart', id, quantity);
    });
});

function charts ( location, labels, data ) {
    var ctx = document.getElementById(location).getContext('2d');
    var categoryCountCharts = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: labels,
            datasets: [{
                data: data,
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
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
            scales: {
                yAxes: [{
                    ticks: {
                        fontColor: 'white',
                        beginAtZero: true
                    }
                }],
                xAxes: [{
                    ticks: {
                        fontColor: 'white',
                    }
                }]
            },
            
        }
    });
}