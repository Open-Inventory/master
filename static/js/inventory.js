$(document).ready(function () {
    var table;

    $("#example").on("mousedown", "td .fas.fa-times", function(e) {
      table.row($(this).closest("tr")).remove().draw();
    })

    $("#example").on('mousedown.edit', "i.far.fa-edit", function(e) {

      $(this).removeClass().addClass("fas fa-check");
      var $row = $(this).closest("tr").off("mousedown");
      var $tds = $row.find("td").not(':last');

      $.each($tds, function(i, el) {
        var txt = $(this).text();
        $(this).html("").append("<input type='text' value=\""+txt+"\">");
      });

    });

    $("#example").on('mousedown', "input", function(e) {
      e.stopPropagation();
    });

    $("#example").on('mousedown.save', "i.fas.fa-check", function(e) {
      
      $(this).removeClass().addClass("far fa-edit");
      var $row = $(this).closest("tr");
      var $tds = $row.find("td").not(':last');
      

      var itemArray = []
      $.each($tds, function(i, el) {
        var txt = $(this).find("input").val()
        itemArray[i] = txt
        $(this).html(txt);
      });

      var item = {
        Id: itemArray[0],
        Name:itemArray[1],
        SubType: itemArray[2],
        Type: itemArray[3],
        ItemStatus: itemArray[4]
      }
      console.log(item)
      var xmlhttp = new XMLHttpRequest();   // new HttpRequest instance 
      xmlhttp.open("POST", "/api/item");
      xmlhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
      xmlhttp.setRequestHeader('X-CSRFToken',document.getElementById('token').getAttribute('value'))
      xmlhttp.send(JSON.stringify(item));
    });
    
    
     $("#example").on('mousedown', "#selectbasic", function(e) {
      e.stopPropagation();
    });
    

    //var testdata = 'static/json/tabledata.json';
    table = $('#example').DataTable({
        stateSave: true,
        ajax: {
          "dataType": 'json',
          "contentType": "application/json; charset=utf-8",
          "type": "GET",
          "url":"/api/inventory",
          "dataSrc": function (json) {
             return json;
          }
      },
      columns: [
        {'data': 'Id'},
        {'data': 'Name'},
        {'data': 'SubType'},
        {'data': 'Type'},
        {'data': 'ItemStatus'},
        {'data': 'Updated'},
        {'data': function(json) {
          return `<i class=\"far fa-edit\" aria-hidden=\"true\"></i> <i class=\"fas fa-times\" aria-hidden=\"true\"></i>`
        }}]
    });
    
    $('#example').css('border-bottom', 'none');

    // add row
    $('#addRow').click(function() {
      var rowHtml = $("#newRow").find("tr")[0].outerHTML
      console.log(rowHtml);
      table.row.add($(rowHtml)).draw();
    });
  });

   