$(document).ready(function () {

    $.getJSON("/api/v1/cars", function (data) {
        dl = $("#carslist")
        dl.children().remove()
        console.log(data)
        $.each(data.data, function (key, val) {
            el = $('<option value="' + val.rn + '"></option>')
            dl.append(el)
        });
    });


    const idElemHistory = "tbodyhistory";
    tbody = document.getElementById(idElemHistory)
    function appendItem(item, isnew) {
        row = document.createElement("tr")
        row.className = "itemhist"
        if (item.oper == "rent") {
            row.className += " return"
        }
        addcolumn = function (value) {
            el = document.createElement("td")
            el.innerHTML = value
            row.appendChild(el)
        }
        addcolumn(item.oper == "rent" ? "Аренда" : "Возврат")
        addcolumn(item.dept)        
        addcolumn(item.agent)
        addcolumn(item.ss)
        addcolumn(item.dateoper)
        tbody.insertBefore(row, tbody.children[0])
        if (isnew) {
            row.className = "newitemhist"
            setTimeout(function (elem, item) {
                elem.className = "itemhist"
                if (item.oper == "rent") {
                    elem.className += " return"
                }
            }, 3000, row, item)
        }
    }

    reload = function () {
        car = $("#carsearch").val();        
        $("#tbodyhistory").children().remove();       
        if (!car) {
            return;
        }        
        $.getJSON("/api/v1/rentjournal/"+car, function (data) {
            $.each(data.data.reverse(), function (key, val) {
                appendItem(val, false);
            });
        });
    }

    ShowError = function (error) {
        var item = document.createElement("div");
        item.innerHTML = "<b> Error: " + error + "</b>";
        item.className = "errorinfo ";
        tab = document.getElementById("thistory");
        document.body.insertBefore(item, tab);   
        setTimeout(function () {
            $(item).hide(2000);
        }, 5000);     
    }

    startWS = function () {
        if (window["WebSocket"]) {
            conn = new WebSocket("ws://" + document.location.host + "/ws");
            conn.onopen = function () {
                $(".errorinfo").remove();
                reload();
            };
            conn.onclose = function (evt) {
                ShowError("WebSocket connection closed. Retry after 5 sec.")
                conn = null
                setTimeout(function () {
                    startWS();
                }, 5000)
            }
            conn.onmessage = function (evt) {
                data = JSON.parse(evt.data);
                car = $("#carsearch").val();      
                if (car && data.rn == car) {
                    appendItem(data, true);
                }
                
            }
        } else {
            ShowError("Error: browser does not support WebSockets")
        }
    }
    
    $("#btn-show").click(function () {        
        event.preventDefault();       
        car = $("#carsearch").val();      
        if (!car) {
            ShowError("Выберете транспортное средство");
        }
        reload();
    });   
    
    $("#btn-clear").click(function(){
        event.preventDefault();   
        $("#tbodyhistory").children().remove(); 
        $("#carsearch").select();
        $("#carsearch").val("");

    });

    $("#carsearch").select();       
    startWS();
});