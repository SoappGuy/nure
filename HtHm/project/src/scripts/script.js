function changeTextSize_onClick() {
    let text =   document.getElementById("text_input");
    let size = document.getElementById("size_input");

    let p = document.getElementById("change_size_p"); 
    console.log(text.textContent);
    p.textContent = text.value;
    p.style.fontSize = size.value + "px";
}

function moveImage_withInterval(){
    let img =   document.getElementById("disappearing_image");
    img.style.left = (Math.random() * window.innerWidth) + "px";
    img.style.top  = (Math.random() * window.innerHeight) + "px";
}

function changeSizeP_onClick() {
    let elements = document.getElementsByTagName("p");
    for(let p of elements) {
        p.setAttribute("style", "font-size:15px");
    }
}

function showTime_onLoad(){
    let date = new Date();
    let h = date.getHours(); // 0 - 23
    let m = date.getMinutes(); // 0 - 59
    let s = date.getSeconds(); // 0 - 59

    h = (h < 10) ? "0" + h : h;
    m = (m < 10) ? "0" + m : m;
    s = (s < 10) ? "0" + s : s;

    let time = h + ":" + m + ":" + s;
    document.getElementById("clock").innerText = time;
    document.getElementById("clock").textContent = time;

    setTimeout(showTime_onLoad, 1000);
}

function vanish() {
    let sections = document.getElementsByTagName("section");

    for(let section of sections) {
        let opacity = 1;
        let timer = setInterval(function() {
            if (opacity <= 0.1) {
                clearInterval(timer);
                section.style.display = 'none';
            }
            section.style.opacity = opacity;
            opacity -= 0.01;
        }, 50);
    }
}

function changeColor_onClick(color) {
    let square = document.getElementById("square");
    square.style.backgroundColor = color;
}

function setMouseStats(e) {
    let mouse_stats = document.getElementById("mouse_stats");

    mouse_stats.textContent = `x: ${e.x} | y: ${e.y}`;
}

function setKeyboardStats(e) {
    let keyboard_stats = document.getElementById("keyboard_stats");

    keyboard_stats.textContent = `key: ${e.code}`
}

function paint_onClick() {
    let colors = ["red", "green", "blue", "yellow", "orange", "brown", "black"];
    let text =   document.getElementById("paintText_input");

    let textDiv = document.getElementById("colored");
    textDiv.innerHTML = '';

    for(let i = 0; i < text.value.length; i++) {
        let span = document.createElement("span");
        span.innerText = text.value[i];
        span.style.color = colors[i % colors.length];
        textDiv.insertAdjacentElement("beforeend", span);
    }
}

function changeFontSize(symbol) {
    let elements = document.querySelectorAll('body *');

    for (let element of elements) {
        let style = window.getComputedStyle(element, null).getPropertyValue('font-size');
        let fontSize = parseFloat(style);
        element.style.fontSize = (fontSize + (symbol == "+" ? 1 : -1)) + 'px';
    }
}

function showInfo(){
    let referrer = document.referrer;
    let user = navigator.userAgent;

    document.getElementById("userInfo").innerHTML = `<p>${referrer}<br>${user}</p>`;
}

function fileInfo(event) {
    const file = event.target.files[0];
    if (!file) {
        return;
    }

    let info = `Розмір файла: ${file.size} байтів<br>`;
    info += `Тип файла: ${file.type}<br>`;
    info += `Остання зміна файла: ${new Date(file.lastModified).toDateString()}<br>`;

    document.getElementById('fileInfo').innerHTML = info;

};

window.setInterval(moveImage_withInterval, 1500);
window.onload = () => {
    showTime_onLoad();
    showInfo();
    document.getElementById('fileInput').onchange = fileInfo;
};

window.addEventListener("mousemove", setMouseStats);
window.addEventListener("keypress", setKeyboardStats);
window.addEventListener("keypress", (e) => {
    if (e.code == "NumpadAdd" || e.code == "Equal") {
        changeFontSize("+");
    } else if (e.code == "NumpadSubtract" || e.code == "Minus") {
        changeFontSize("-");
    }
});