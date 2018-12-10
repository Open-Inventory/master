
function format ( productArray ) {
  let productRow='<table cellpadding="5" cellspacing="0" border="0" style="padding-left:500x;"><tr>';
  for(var i=0; i < productArray.length; i++) {
     productRow=productRow+`<tr><td>Product ${i+1}:</td><td>${productArray[i].Products}</td><td>ID: ${productArray[i].ItemID}</td>`;
  }
  return productRow+'</table>';
}


$(document).ready(function () {
  var table;

  $("#example").on("mousedown", "td .fas.fa-times", function (e) {
    table.row($(this).closest("tr")).remove().draw();
  })

  $('#example').on('mousedown', "td .fas.fa-plus", function () {
    $(this).removeClass().addClass("fas fa-minus");    
    var tr = $(this).closest('tr');
    var row = table.row( tr );

    // Get the order quantities for each order
    var size = table.rows().count();
    var length=orderList.length;
    var orderQuantities = [];
    for (var i = 1; i <= orderList[length-1].ID; i++) {
      orderQuantities[i] = 0;
      for (var j = 0; j < length; j++) {

        if (orderList[j].ID == i) {
          orderQuantities[i]++;
        }
      }
    } //OrderQuantities[orderID] now holds the quantity for each order.
    var quantityByItem = [];
    for (var i = 1; i <= orderQuantities.length; i++) {
      //quantityByItem[i] = 0;
      for (var j = 0; j < length; j++) {
        if (orderList[j].ID == i) {
          quantityByItem[j] = orderQuantities[i];
        }
      }
    }

    

    //get current row index
    for (var i = 0; i < size; i++) { 
      if (table.row(i).data().ItemID == row.data().ItemID) {
        // console.log(i, " ", table.row(i).data());
        // console.log(row.data());
        currentIndex = i;
      }
    }
    var nextrow = table.row ( currentIndex + 1 );

    // console.log("\nCurrent Row:",row.data());
    // console.log("Next Row:",nextrow.data());

    var productArray = []
    for (var i = 0; i < quantityByItem[currentIndex]; i++) { //row.data().Quantity
      productArray.push(table.row(currentIndex + i).data());
    }
    // console.log("i:", i, "\nindex:", currentIndex, "\nquantity:", quantityByItem[currentIndex]);
    // console.log(productArray);

    // Open this row
    row.child( format(productArray) ).show();
    tr.addClass('shown');

  });

  $('#example').on('mousedown', "td .fas.fa-minus", function () {
    $(this).removeClass().addClass("fas fa-plus");
  
    var tr = $(this).closest('tr');
    var row = table.row( tr );

    // This row is already open, close it
    row.child.hide();
    tr.removeClass('shown');
  });

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
      ItemID: orderArray[3],
      Status: orderArray[4],
    }

    // $.getJSON("/api/item?name=" + orderArray[0] + "&Products=" + item.Products + "&Status=" + order.Status, function (response) {
      $.getJSON("/api/item?name=" + itemArray[0] + "&subtype=" + itemArray[1] + "&type=" + itemArray[2] + "&status=" + itemArray[3], function (response) {

    });
  });


  $("#example").on('mousedown', "#selectbasic", function (e) {
    e.stopPropagation();
  });

  // var productArray = [];
  var count = 0;
  var orderList = [];

  //var testdata = 'static/json/tabledata.json';
  table = $('#example').DataTable({

  "rowCallback": function( row, data ) {
    var multipleProducts = 0;
    var dTable = table.data();
    orderList.push(dTable[count]);
    if(count != 0) {
      if (dTable[count].ID == dTable[count-1].ID) {
        multipleProducts = 1;
        $(row).hide();
      }
    }
    count++;
  },
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
      // { 'data': 'Quantity' },
      // { 'data': 'Products' }, 
      { 'data': 'Status' },
      { 'data': 'Date' },
      {
        'data': function (json) {
          return `<i class=\"fas fa-plus\" aria-hidden=\"true\"></i> <i class=\"far fa-edit\" aria-hidden=\"true\"></i> <i class=\"fas fa-times\" aria-hidden=\"true\"></i>`
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

