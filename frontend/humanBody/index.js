window.onload = function () {
    
    const pieces = document.getElementsByTagName('svg');
    for (var i = 0; pieces.length; i++) {
        let _piece = pieces[i];
        _piece.onclick = function(t) {
            if (t.target.getAttribute('data-position') != null) document.getElementById('data').innerHTML = t.target.getAttribute('data-position');
            if (t.target.parentElement.getAttribute('data-position') != null) document.getElementById('data').innerHTML = t.target.parentElement.getAttribute('data-position');
        }
    }
}

// createBox = function (){
//     const newDiv = document.createElement("div");
//     const newContent = document.createTextNode("example1");
//     newDiv.appendChild(newContenct);
const card_template = document.querySelector("#exercise-card-template");
const exercise_list = document.querySelector(".exercises");

fetch("http://localhost:4000/v1/exercises")
    .then(res => res.json())
    .then(data => {
        data.forEach(exercise => {
            const card = card_template.content.cloneNode(true).children[0];
            const exercise_name = card.querySelector("[exercise-name]");
            const thumbnail = card.querySelector(".thumbnail");
            exercise_name.innerHTML= exercise.Name;
            thumbnail.src = exercise.URL;
            console.log(card);
            exercise_list.append(card);
        })
    });
