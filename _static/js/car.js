$(document).ready(function () {

   var selectedCarInfo;

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

    reloadcar = function(car) {
        selectedCarInfo = null;
        $.getJSON("/api/v1/car/"+car, function (data) {            
            if (data.error) {
                alert(data.error);
                return;
            }
            selectedCarInfo  = data.data;
            $("#car-rn").text(data.data.rn);
            $("#car-type").text(data.data.model.cartype.type);
            $("#car-model").text(data.data.model.name);                      
            isrent = data.data.isrent === 1
            $("#agent").prop('disabled',isrent);
            if (isrent) {
                $("#car-status").text("В аренде. Арендатор:"+data.data.agent
                    +".  Взято в:"+data.data.department+"  ("+data.data.dateoper+")").removeClass("car-goods").addClass("car-rent");
                $("#caraction").text("Вернуть");
                $("#agent").val(data.data.agent);
                $("#dep").val("");

            } else {
                $("#car-status").text("В наличии в:"+data.data.department).removeClass("car-rent").addClass("car-goods");
                $("#caraction").text("Взять в аренду");
                $("#agent").val("");
                $("#dep").val("");
            }         
            showcar();
        })
            .fail(function(){
                ShowError("Транспортное средство "+car+" не найдено");               
                clear();
                $("#carsearch").val(car);
            });
    } 
    reload = function (car) {                
        $("#tbodyhistory").children().remove();       
        if (!car) {
            return;
        }        
        $.getJSON("/api/v1/rentjournal/"+car, function (data) {
            $.each(data.data.reverse(), function (key, val) {
                appendItem(val, false);
            });
        }); 
        reloadcar(car);       
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
    
    showcar = function(){        
        $("#carinfo").removeClass("hide").addClass("show")        
        $("#carsearch").prop('disabled',true);
        $("#btn-show").prop('disabled',true);
    };    
    
    clear = function(){
        $("#tbodyhistory").children().remove(); 
        $("#carsearch").prop('disabled',false);
        $("#carsearch").select();
        $("#carsearch").val("");
        $("#carinfo").removeClass("show").addClass("hide")
        $("#btn-show").prop('disabled',false);
    }

    $("#btn-show").click(function () {        
        event.preventDefault();  
        car = $("#carsearch").val();
        if (!car) {
            ShowError("Выберете транспортное средство");
            return;
        }        
        reload(car);        
    });  
    $("#btn-clear").click(function(){
        event.preventDefault();   
        clear();
    });

    rent = function(prn,pdep,pagent) {
        //TODO: определить что аренда или возврат
        var apiurl = "/api/v1/rent";
        if (selectedCarInfo.isrent===1){
            apiurl = "/api/v1/return";
        }
        $.post(apiurl,{regnum:prn,dept:pdep,agent:pagent})
            .done(function(){
                //reloadcar();
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
            alert("Выберите ТС");
            car.select();
            return;
        }              
        if (dep.val()==="") {
            alert("Выберите подразделение");
            dep.select();
            return;
        }
        if (agn.val()==="") {
            alert("Введите ФИО");
            agn.select();
            return;
        }
        rent(car.val(),dep.val(),agn.val());
    });


    reload();
    $.getJSON("/api/v1/cars", function (data) {
        dl = $("#carslist")
        dl.children().remove()        
        $.each(data.data, function (key, val) {
            el = $('<option value="' + val.rn + '"></option>')
            dl.append(el)
        });
    });

    $("#carsearch").select();       
   
    if ($("#carsearch").val()!=="") {
        showcar();
    }

    mustreload = false;
    wsopen=function(){
        $(".errorinfo").remove();
        if (mustreload) {
            // Для загрузки после переподключения 
            reload();                    
        }  
    }
    wsclose=function(evt) {
        mustreload = true;
        ShowError("WebSocket connection closed. Retry after 5 sec.")
    }
    wsmessage=function (evt) {
        data = JSON.parse(evt.data);
        car = $("#carsearch").val();      
        if (car && data.rn == car) {
            appendItem(data, true);
        }
        reloadcar(car);  
    }
    if (!window.startws(5000,wsmessage,wsopen,wsclose)) {
        ShowError("Error: browser does not support WebSockets")
    }

});