window.onload=function() {
    console.log("test")    
    const idElemHistory = "tbodyhistory";

    tbody = document.getElementById(idElemHistory)
    function appendItem(item,isnew) {
        row = document.createElement("tr")
        row.className = "itemhist"
        if (item.oper == "rent") {
            row.className += " return"
        }
        addcolumn = function(value) {
            el = document.createElement("td")
            el.innerHTML =  value
            row.appendChild(el)
        } 
        addcolumn(item.type)              
        addcolumn(item.model)
        addcolumn( '<a href="/carhistory">'+item.rn+'<a>')
        addcolumn(item.agent)
        addcolumn(item.ss)
        addcolumn(item.oper=="rent"?"Аренда":"Возврат" )
        addcolumn(item.dateoper)    
      
        tbody.insertBefore(row,tbody.children[0])
        if (isnew) {
            row.className = "newitemhist"
        
            setTimeout(function(elem,item){
                elem.className = "itemhist"
                if (item.oper == "rent") {
                    elem.className += " return"
                }
            },3000,row,item)
        }
    }

    $.getJSON( "/api/v1/rentjournal", function( data ) {               
        $.each( data.data.reverse(), function( key,val ) {
            appendItem(val,false)
        });              
    });


    ShowError = function (error) {
        var item = document.createElement("div");
        item.innerHTML = "<b> Error: "+error+"</b>";
        item.className = "errorinfo "
        
        tab = document.getElementById("thistory")
        document.body.insertBefore(item,tab)
        
    }

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://"+ document.location.host+"/ws");
        conn.onclose = function(evt) {
            ShowError("WebSocket connection closed.")            
        }
        conn.onmessage = function(evt) {
            appendItem( JSON.parse(evt.data),true)
        }
    } else {
        ShowError("Error: browser does not support WebSockets")
    }

    
}