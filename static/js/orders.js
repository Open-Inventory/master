$(document).ready(function () {
  var table;

  $("#example").on("mousedown", "td .fas.fa-times", function (e) {
    table.row($(this).closest("tr")).remove().draw();
  })

  $("#example").on('mousedown.edit', "i.far.fa-edit", function (e) {

    $(this).removeClass().addClass("fas fa-check");
    var $row = $(this).closest("tr").off("mousedown");
    var $tds = $row.find("td").not(':last');

    $.each($tds, function (i, el) {
      var txt = $(this).text();
      $(this).html("").append("<input type='text' value=\"" + txt + "\">");
    });

  });

  $("#example").on('mousedown', "input", function (e) {
    e.stopPropagation();
  });

  $("#example").on('mousedown.save', "i.fas.fa-check", function (e) {

    $(this).removeClass().addClass("far fa-edit");
    var $row = $(this).closest("tr");
    var $tds = $row.find("td").not(':last');


    var orderArray = []
    $.each($tds, function (i, el) {
      var txt = $(this).find("input").val()
      orderArray[i] = txt
      $(this).html(txt);
    });

    var order = {
      ID: orderArray[0],
      Quantity: orderArray[1],
      Products: orderArray[2],
      Status: orderArray[3],
    }

    // $.getJSON("/api/item?name=" + orderArray[0] + "&Products=" + item.Products + "&Status=" + order.Status, function (response) {
      $.getJSON("/api/item?name=" + itemArray[0] + "&subtype=" + itemArray[1] + "&type=" + itemArray[2] + "&status=" + itemArray[3], function (response) {

    });
  });


  $("#example").on('mousedown', "#selectbasic", function (e) {
    e.stopPropagation();
  });


  //var testdata = 'static/json/tabledata.json';
  table = $('#example').DataTable({
    stateSave: true,
    ajax: {
      "dataType": 'json',
      "contentType": "application/json; charset=utf-8",
      "type": "GET",
      "url": "/api/orders",
      "dataSrc": function (json) {
        return json;
      }
    },
    columns: [
      { 'data': 'ID' },
      { 'data': 'Quantity' },
      // { 'data': 'Products' }, 
      { 'data': 'Status' },
      { 'data': 'Date' },
      {
        'data': function (json) {
          return `<i class=\"far fa-edit\" aria-hidden=\"true\"></i> <i class=\"fas fa-times\" aria-hidden=\"true\"></i> <i class=\"fas fa-plus\" aria-hidden=\"true\"></i>`
        }
      }]
  });

  $('#example').css('border-bottom', 'none');

  // add row
  $('#addRow').click(function () {
    var rowHtml = $("#newRow").find("tr")[0].outerHTML
    console.log(rowHtml);
    table.row.add($(rowHtml)).draw();
  });
});

