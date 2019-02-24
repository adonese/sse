function onLoaded() {
    var source = new EventSource("/sse");
    source.onmessage = function (event) {
        console.log("on message please kill the fuck off me!");
        console.dir(event);
        document.getElementById("counter").innerHTML = event.data;
    }
}