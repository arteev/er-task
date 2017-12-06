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

    $.getJSON("/api/v1/departments", function (data) {
        dl = $("#departmentslist")
        $.each(data.data, function (key, val) {
            el = $('<option value="' + val.name + '"></option>')
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
        })            
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

    
    
    showcar = function(){
        car = $("#carsearch").val();      
        if (!car) {
            ShowError("Выберете транспортное средство");
        } else {
            $("#carinfo").removeClass("hide").addClass("show")
        }
       
        $("#carsearch").prop('disabled',true);
        $("#btn-show").prop('disabled',true);
    };

    $("#btn-show").click(function () {        
        event.preventDefault();       
        showcar();
        reload();
    });   
    
    $("#btn-clear").click(function(){
        event.preventDefault();   
        $("#tbodyhistory").children().remove(); 
        $("#carsearch").prop('disabled',false);
        $("#carsearch").select();
        $("#carsearch").val("");
        $("#carinfo").removeClass("show").addClass("hide")
        $("#btn-show").prop('disabled',false);
    });

    rent = function(prn,pdep,pagent) {
        //TODO: определить что аренда или возврат
        $.post("/api/v1/rent",{regnum:prn,dept:pdep,agent:pagent})
            .done(function(){
                alert("ТС Взято в аренду");
            })
            .fail(function(data){
                console.log(data);
                if (data.responseJSON.error) {
                    alert(data.responseJSON.error);
                } else {
                    alert("Произошла ошибка: " + data.status +" - "+ data.statusText);
                }                                 
            });        

    };

    $("#caraction").click(function(){
        car = $("#carsearch");
        dep = $("#dep");
        agn = $("#agent");
        if (car.val()==="") {
            alert("Выберите ТС")
            car.select();
            return;
        }              
        if (dep.val()==="") {
            alert("Выберите подразделение")
            dep.select();
            return;
        }
        if (agn.val()==="") {
            alert("Введите ФИО")
            agn.select();
            return;
        }
        rent(car.val(),dep.val(),agn.val());
    });

    $("#carsearch").select();       
    startWS();
    if ($("#carsearch").val()!=="") {
        showcar();
    }
});