// блок 1:

// 1. Створити функцію, яка виводить текст з різним розміром шрифту.
//    Функція має два аргументи: текстовий рядок і розмір шрифту.
//    Використовуйте style-властивість fontSize
function textSize() {
    let text =   document.getElementById("text1");
    let size = document.getElementById("size1");
    changeSize(text.value, size.value);
}

function changeSize(text, size) {
    let p = document.getElementById("change_size");
    p.textContent = text;
    p.style.fontSize = size + "px";
}

// 2. Використовуючи властивості елемента IMG style.top і style.left,
//    зробіть сторінку, на якій невелика картинка кожну секунду виникає в новому
//    місці екрану. Використовуйте setInterval.
function moveImage() {
    let img =   document.getElementById("imageMove");
    img.style.left = (Math.random() * window.innerWidth) + "px";
    img.style.top  = (Math.random() * window.innerHeight) + "px";
}

window.setInterval(moveImage, 1500);

// 3. Знайти на сторінці всі <p> і змінити їх розмір на 15px.
//    Використовувати getElementsByTagName, setAttribute
function changeP() {
    let elements = document.getElementsByTagName("p");
    for(let p of elements) {
        p.setAttribute("style", "font-size:15px");
    }
}

// 4. Текстовий годинник – використовувати функції setInterval або
//    setTimeout об'єкта window.
function clock(){
    let date = new Date();
    let h = date.getHours();
    let m = date.getMinutes();
    let s = date.getSeconds();

    h = (h < 10) ? "0" + h : h;
    m = (m < 10) ? "0" + m : m;
    s = (s < 10) ? "0" + s : s;

    let time = h + ":" + m + ":" + s;
    document.getElementById("clock").innerText = time;
    document.getElementById("clock").textContent = time;

    setTimeout(clock, 1000);
}

// 5. Створити ефект поступового витирання (аналог фільтра) частини
//    документа, використовуючи таймер.
function clearScreen() {
    let blocks = document.getElementsByClassName("block");

    for(let block of blocks) {
        let opacity = 1;
        let timer = setInterval(
            () => {
                block.style.opacity = opacity;
                opacity -= 0.1;
            },
            50);
    }
}

// блок 2:

// 3. У лівому верхньому кутку екрану - список з п'яти кольорів. У правому
//    верхньому куті - чорний квадрат. Колір квадрата можна поміняти,
//    використовуючи список.
function squareColor(color) {
    let square = document.getElementById("square");
    square.style.backgroundColor = color;
}

// 4. Створіть сторінку, на якій в DIV-блок будуть виводиться поточні
//    координати курсору мишки і код натиснутої клавіші на клавіатурі.
function mouseStats(e) {
    let mouse_stats = document.getElementById("mouse_stats");

    mouse_stats.textContent = `x: ${e.x}  y: ${e.y}`;
}

function keyboardStats(e) {
    let keyboard_stats = document.getElementById("keyboard_stats");

    keyboard_stats.textContent = `key: ${e.code}`
}

window.addEventListener("mousemove", mouseStats);
window.addEventListener("keypress", keyboardStats);

// 5. Создать сторінку, розмір тексту на якій можна змінювати за
//    допомогою кнопок A + / A-. Вибір користувача запомнінать в cookie і при
//    наступному вході на сторінку відразу відображати її запомненним розміром
//    шрифту.
function restoreFontSize() {
    const match = document.cookie.match(new RegExp(`(^| )fontSize=([^;]+)`))
    if (match) {
        let fontSize = match[2];

        let element = document.getElementById('cookie_text');
        element.style.fontSize = fontSize + "px";
    }
}

function cookieSize(e) {
    function changeFontSize(size) {
        let element = document.getElementById('cookie_text');

        let fontSize = parseFloat(window.getComputedStyle(element, null).getPropertyValue('font-size'));

        let newFontSize = (fontSize + size);
        element.style.fontSize = newFontSize + 'px';

        let expires = "";
        let date = new Date();
        date.setTime(date.getTime() + (7*24*60*60*1000));
        expires = "; expires=" + date.toUTCString();

        document.cookie = "fontSize" + "=" + (newFontSize || "")  + expires + "; path=/";
        console.log(document.cookie);
    }

    if (e.code === "NumpadAdd" || e.code === "Equal") {
        changeFontSize(1);
    } else if (e.code === "NumpadSubtract" || e.code === "Minus") {
        changeFontSize(-1);
    }
}

window.addEventListener("keypress", cookieSize);

// 6. Вивести на екран фразу, розфарбувавши її букви в заданий діапазон
//    кольорів (якщо букв більше, ніж кольорів, то застосовувати кольори по колу).
function colorizeText() {
    let colors = ["black", "red", "green", "blue", "yellow", "white"];
    let text =   document.getElementById("text2");

    let textDiv = document.getElementById("colored_text");
    textDiv.innerHTML = '';

    for(let i = 0; i < text.value.length; i++) {
        let span = document.createElement("span");
        span.innerText = text.value[i];
        span.style.color = colors[i % colors.length];
        textDiv.insertAdjacentElement("beforeend", span);
    }
}


// блок 3:
// 5. Звідки ви прийшли. Версія вашого браузера.
function userInfo(){
    let referrer = document.referrer;
    let userAgent = navigator.userAgent;

    document.getElementById("user_info").innerHTML = `${referrer}<br>${userAgent}`;
}

// 8. Реалізувати скрипт знаходження інформації про фото: дата
// створення/зміни, розмір, власник.
function fileInfo(event) {
    const file = event.target.files[0];
    if (!file) {
        return;
    }

    let info = `Розмір файла: ${file.size} байтів<br>`;
    info += `Остання зміна: ${new Date(file.lastModified).toDateString()}<br>`;
    info += `Тип файла: ${file.type}<br>`;

    document.getElementById('file_info').innerHTML = info;
}