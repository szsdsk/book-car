const body = document.querySelector('body'),
    form = document.querySelector('form'),
    fieldset = form.querySelectorAll('fieldset'),
    count = fieldset.length;
const map = new Map();
const index = new Map();
const message = document.querySelector('.el-message-box__wrapper');
const wrapper = document.querySelector('.v-modal');

// 点击确定返回主页
document.querySelector('.el-button--primary').addEventListener('click', () => {
    message.style.display = 'none';
    wrapper.style.display = 'none';
    window.location.href = '/index';
});

// 渲染表格
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

// 点击相应的点跳到某一项
document.querySelector('ul.items').addEventListener('click', event => {
    let target = event.target.closest('li');
    let active = document.querySelector("ul.items li.active");
    if (target === active) {
        return;
    }
    target.classList.add('active');
    active.classList.remove('active');
    map.get(target).classList.add('enable');
    map.get(active).classList.remove('enable');
    let input = map.get(target).querySelector("input");
    if (input !== null) {
        input.focus();
    }
});

// 鼠标滚轮滑动切换数据
document.addEventListener('wheel', event => {
    let active = document.querySelector("ul.items li.active");
    if (event.deltaY < 0 && index.get(active) > 0) {
        active.classList.remove('active');
        map.get(active).classList.remove('enable');
        active.previousElementSibling.classList.add('active');
        map.get(active).previousElementSibling.classList.add('enable');
        let input = map.get(active).previousElementSibling.querySelector("input");
        if (input !== null) {
            input.focus();
        }
    } else if (event.deltaY > 0 && index.get(active) < count - 1) {
        active.classList.remove('active');
        map.get(active).classList.remove('enable');
        active.nextElementSibling.classList.add('active');
        map.get(active).nextElementSibling.classList.add('enable');
        let input = map.get(active).nextElementSibling.querySelector("input");
        if (input !== null) {
            input.focus();
        }
    }

});

window.onload = init;

// 进入下一项
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
        nextEnable.classList.add("enable");
        let input = nextEnable.querySelector("input");
        if (input !== null) {
            input.focus();
        }
        // Switch active class on left list
        let active = document.querySelector("ul.items li.active"),
            nextActive = active.nextElementSibling;
        active.classList.remove("active");
        nextActive.classList.add("active");
    }
}
// 回车也可以到下一项
form.addEventListener("keydown", keyDown, false);
function keyDown(event) {
    let key = event.key,
        target = document.querySelector("fieldset.enable .button");
    if (key === 'Enter') {
        event.preventDefault();
        next(target);
    }
}
// 点击按钮也可以到下一项
form.addEventListener('click', function(event) {
    let target = event.target.closest('div');
    if (target.classList.contains("button"))
        next(target);
});

// 发送请求，用于新创建记录
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
    if (pickUpTime > dropOfTime) {
        alert('还车时间小于取车时间')
        return
    }
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
    let response = await fetch('/booking/submit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
    });
    let json = await response.json();
    message.style.display = 'flex';
    wrapper.style.display = 'block';
    document.querySelector('.price').innerText =" 此次预订的价格为：" + json.data.price + "。 点击确定返回首页。"
});