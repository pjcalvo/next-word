let seconds =  document.getElementById("initialTime").innerHTML
if (seconds){

    var x = setInterval(function() {

    // Reduce the timer
    seconds = seconds - 1;

    // Display the result in the element with id="demo"
    document.getElementById("counter").innerHTML = `<strong>${seconds}</strong> segundos`

    // If the count down is finished, write some text
    if (seconds <= 0) {
        clearInterval(x);
        document.getElementById("counter").innerHTML = "TIEMPO AGOTADO";
    }
    }, 1000);
} 