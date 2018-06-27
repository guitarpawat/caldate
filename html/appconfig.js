$(document).ready(function(){
    $("button").click(function(){

        var showData = $('#output');
        //alert("65555");
        
        $.getJSON("example.json", function (data) {
          alert(1);
          // var items = data.items.map(function (item) {
          //   return item.key + ': ' + item.value;
          // });
    
          // showData.empty();
    
          // if (items.length) {
          //   var content = '<li>' + items.join('</li><li>') + '</li>';
          //   var list = $('<ul />').html(content);
          //   showData.append(list);
          //   alert(2);
          // }
        });
    
        showData.text('Loading the JSON file.');
    });
});