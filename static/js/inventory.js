$(document).ready(function () {
    var table;

    $("#example").on("mousedown", "td .fas.fa-times", function(e) {
      table.row($(this).closest("tr")).remove().draw();
    })

    $("#example").on('mousedown.edit', "i.far.fa-edit", function(e) {

      $(this).removeClass().addClass("fas fa-check");
      var $row = $(this).closest("tr").off("mousedown");
      var $tds = $row.find("td").not(':last');
      var $tds = $tds.not(':last');

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
      
      $.each($tds, function(i, el) {
        var txt = $(this).find("input").val()
        $(this).html(txt);
      });
    });
    
    
     $("#example").on('mousedown', "#selectbasic", function(e) {
      e.stopPropagation();
    });
    

    var testdata = 'static/json/tabledata.json';
    table = $('#example').DataTable({
        stateSave: true,
        ajax: testdata,
        columns: [{
        data: 'item'
        }, {
        data: 'manufacturer'
        }, {
        data: 'category'
        }, {
        data: 'location'
        }, {
        data: 'stock'
        }, {
        data: 'action'
        }]
    });
    
    $('#example').css('border-bottom', 'none');

    // add row
    $('#addRow').click(function() {
      var rowHtml = $("#newRow").find("tr")[0].outerHTML
      console.log(rowHtml);
      table.row.add($(rowHtml)).draw();
    });
  });

   