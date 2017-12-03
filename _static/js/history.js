$(document).ready(function () {

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
        addcolumn(item.type)
        addcolumn(item.model)
        addcolumn('<a href="/car/' + item.rn + '">' + item.rn + '<a>')
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
        $("#tbodyhistory").children().remove()
        $.getJSON("/api/v1/rentjournal", function (data) {
            $.each(data.data.reverse(), function (key, val) {
                appendItem(val, false)
            });
        });
    }


    ShowError = function (error) {
        var item = document.createElement("div");
        item.innerHTML = "<b> Error: " + error + "</b>";
        item.className = "errorinfo "
        tab = document.getElementById("thistory")
        document.body.insertBefore(item, tab)
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
                appendItem(JSON.parse(evt.data), true);
            }
        } else {
            ShowError("Error: browser does not support WebSockets")
        }
    }


    startWS()

    $("#btn-refresh").click(function () {
        reload();
    })

    $("#btn-car").click(function () {
        window.location = "/car"
    })


});