

const form = document.querySelector("form")
const input = document.querySelector("#get-num")
function addEvent() {
    const buttons = document.querySelectorAll(".btn--full")
    buttons.forEach(button => {
        button.addEventListener('click', () => {
            window.location.href = `/booking/${button.id}`
        });
    });
}

addEvent();
form.addEventListener("submit", async function(ev) {
    ev.preventDefault();
    const size = +input.value;
    let response = await fetch(`/index/${size}`)
    let data = await response.json();
    const cardContainer = document.querySelector('.support-card-container');
    cardContainer.innerHTML = '';
    data.Cars.forEach(car => {
        // 创建支持卡片元素
        const supportCard = document.createElement('div');
        supportCard.classList.add('support-card', 'grid', 'grid-2-col');

        // 创建图像盒子
        const imageBox = document.createElement('div');
        imageBox.classList.add('image-box', 'flex');

        // 创建图片元素
        const image = document.createElement('img');
        image.classList.add('card-image');
        image.src = car.img;
        image.alt = car.model;

        // 将图片元素添加到图像盒子中
        imageBox.appendChild(image);

        // 创建文本盒子
        const textBox = document.createElement('div');
        textBox.classList.add('card-text');

        // 创建标题元素
        const heading = document.createElement('h3');
        heading.classList.add('heading-third');
        heading.textContent = car.Model;

        // 创建座位数元素
        const seats = document.createElement('p');
        seats.textContent = `Seats: ${car.capacity}`;

        // 创建描述元素
        const description = document.createElement('p');
        description.textContent = car.description;

        // 创建价格元素
        const price = document.createElement('p');
        price.innerHTML = `Price: $${car.pricePerHour.toFixed(1)} per&nbsp;hour &nbsp; $${car.pricePerDay.toFixed(1)} per&nbsp;day`;

        // 创建按钮元素
        const button = document.createElement('button');
        button.id = car.Id;
        button.classList.add('btn', 'btn--full');
        button.textContent = 'Book Now';

        // 将所有元素添加到文本盒子中
        textBox.append(heading);
        textBox.append(seats);
        textBox.append(description);
        textBox.append(price);
        textBox.append(button);

        // 将图像盒子和文本盒子添加到支持卡片中
        supportCard.append(imageBox);
        supportCard.append(textBox);

        // 将支持卡片添加到容器中
        cardContainer.append(supportCard);
    })
    addEvent();
})

