const body = document.querySelector('body'),
    form = document.querySelector('form'),
    fieldset = form.querySelectorAll('fieldset'),
    count = fieldset.length;
const map = new Map();
const index = new Map();
function init() {
    // 创建fieldset等量的li
    const ul = document.querySelector('ul.items');
    for (let i = 0; i < count; i++) {
        let li = document.createElement('li');
        map.set(li, fieldset[i]);
        index.set(li, i);
        ul.append(li);
    }
    // Add class active on first li
    ul.firstElementChild.classList.add('active');
}

document.querySelector('ul.items').addEventListener('click', event => {
    let target = event.target.closest('li');
    let active = document.querySelector("ul.items li.active");
    target.classList.add('active');
    active.classList.remove('active');
    map.get(target).classList.add('enable');
    map.get(active).classList.remove('enable');
});

document.addEventListener('wheel', event => {
    let active = document.querySelector("ul.items li.active");
    if (event.deltaY < 0 && index.get(active) > 0) {
        active.classList.remove('active');
        map.get(active).classList.remove('enable');
        active.previousElementSibling.classList.add('active');
        map.get(active).previousElementSibling.classList.add('enable');
    } else if (event.deltaY > 0 && index.get(active) < count - 1) {
        active.classList.remove('active');
        map.get(active).classList.remove('enable');
        active.nextElementSibling.classList.add('active');
        map.get(active).nextElementSibling.classList.add('enable');
    }

});

window.onload = init;

function next(target) {
    let input = target.previousElementSibling;
    // Check if input is empty
    if (input.value === "") {
        body.classList.add("error");
    } else {
        body.classList.remove("error");
        let enable = document.querySelector("form fieldset.enable"),
            nextEnable = enable.nextElementSibling;
        enable.classList.remove("enable");
        enable.classList.add("disable");
        nextEnable.classList.add("enable");

        // Switch active class on left list
        let active = document.querySelector("ul.items li.active"),
            nextActive = active.nextElementSibling;
        active.classList.remove("active");
        nextActive.classList.add("active");
    }
}

form.addEventListener("keydown", keyDown, false);
function keyDown(event) {
    let key = event.key,
        target = document.querySelector("fieldset.enable .button");
    if (key === 'Enter') {
        event.preventDefault();
        next(target);
    }
}
form.addEventListener('click', function(event) {
    let target = event.target.closest('div');
    if (target.classList.contains("button"))
        next(target);
});

form.addEventListener("submit", async function(event) {
    event.preventDefault();
    const uid = form.elements["userName"].value;
    const firstName = form.elements["firstName"].value;
    const lastName = form.elements["lastName"].value;
    const creditCard = form.elements["creditcard"].value;
    const phone = form.elements["phone"].value;
    const telephone = form.elements["telephone"].value;
    const licence = Number(form.elements["licence"].value);
    const stateIssue = form.elements["stateissue"].value;
    const isStudent = form.elements["job"].value === "student";
    const tickets = Number(form.elements["tickets"].value);
    const expiration = new Date(form.elements["expiration"].value);
    const carId = Number(form.elements["carID"].value);
    const email = form.elements["email"].value;
    const address = form.elements["address"].value;
    const locationId = Number(form.elements["locationId"].value);
    const pickUpTime = new Date(form.elements["startTime"].value);
    const dropOfTime = new Date(form.elements["endTime"].value);
    const formData = {
        uid,
        firstName,
        lastName,
        creditCard,
        phone,
        telephone,
        licence,
        stateIssue,
        isStudent,
        tickets,
        expiration,
        pickUpTime,
        dropOfTime,
        carId,
        email,
        address,
        locationId
    };
    await fetch('/booking/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
    });
    // window.location.href = '/index';
});