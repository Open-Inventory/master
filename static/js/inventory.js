var t

$(document).ready(function () {
    /*
    $('#table_id tbody td').editable( function( sValue ) {
		//Get the position of the current data from the node 
		var aPos = t.api().fnGetPosition( this );
		
		//Get the data array for this row 
		var aData = t.fnGetData( aPos[0] );
		
		//Update the data array and return the value 
		aData[ aPos[1] ] = sValue;
		return sValue;
    }, { "onblur": 'submit' } ); // Submit the form when bluring a field 
    */
    
    t = $('#table_id').DataTable();

    $('#AddRow').on( 'click', function () {
        t.api().row.add( [
            "Click to edit",
            "Click to edit",
            "Click to edit",
            "Click to edit",
            "Click to edit"
        ] ).draw( false );
     });

     $('#table_id tbody').on( 'click', 'tr', function () {
        if ( $(this).hasClass('selected') ) {
            $(this).removeClass('selected');
        }
        else {
            t.api().$('tr.selected').removeClass('selected');
            $(this).addClass('selected');
        }
    });

 
    $('#RemoveRow').click( function () {
        t.api().row('.selected').remove().draw( false );
    } );
});