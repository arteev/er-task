window.onload=function() {
    console.log("test")    
    const idElemHistory = "tbodyhistory";

    tbody = document.getElementById(idElemHistory)
    function appendItem(item) {
        row = document.createElement("tr")
        row.className = "newitemhist"
        addcolumn = function(value) {
            el = document.createElement("td")
            el.innerHTML =  value
            row.appendChild(el)
        } 
        addcolumn(item.type)              
        addcolumn(item.model)
        addcolumn( '<a href="/carhistory">'+item.rn+'<a>')
        addcolumn(item.daterent)
        addcolumn(item.dateret)
        addcolumn(item.agent)
        tbody.insertBefore(row,tbody.children[0])
        setTimeout(function(elem){
            elem.className = "itemhist"
        },3000,row)
    }

    ShowError = function (error) {
        var item = document.createElement("div");
        item.innerHTML = "<b> Error: "+error+"</b>";
        item.className = "errorinfo"
        
        tab = document.getElementById("thistory")
        document.body.insertBefore(item,tab)
        
    }

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://"+ document.location.host+"/ws");
        conn.onclose = function(evt) {
            ShowError("WebSocket connection closed.")            
        }
        conn.onmessage = function(evt) {
            appendItem( JSON.parse(evt.data))
        }
    } else {
        ShowError("Error: browser does not support WebSockets")
    }

    
}