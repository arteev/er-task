$(document).ready(function () {
    getStats =function(url,id) {
        $.getJSON(url, function (data) {
            tabbody = $(id)      
            tabbody.children().remove();
            if (data.data) {            
                $.each(data.data, function (key, val) {
                   row =$("<tr></tr>");
                   row.append($('<td colspan="3">'+val.department+'</td>')).addClass("table-row-group");
                   tabbody.append(row);        
                   var ln=val.entities.length;     
                   $.each(val.entities, function (key, entity) {
                    row =$("<tr></tr>");
                    
                    if (ln>0 ){
                        row.append($('<td rowspan="'+ln+'"></td>'));
                        ln = 0;
                    }

                    row.append($('<td>'+entity.entity+'</td>'));
                    row.append($('<td>'+entity.avgduration+'</td>'));               
                    tabbody.append(row);                
                   });               
                });
            }
        });
    }
    window.reloadstats = function() {
        getStats("/api/v1/stats/deps/model","#stats-model-dep-body");
        getStats("/api/v1/stats/deps/type","#stats-type-dep-body"); 
    }
    
    window.reloadstats();
});