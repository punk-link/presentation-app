(() => {
    "use strict";
    const links = document.querySelectorAll(".album__link");
    const dots = document.querySelector(".dots");
    if (links.length > 1) dots.style.display = "block";
    links.forEach((function(e, i) {
        if (i > 0) e.classList.add("hidden");
    }));
    dots.addEventListener("click", (function() {
        links.forEach((function(e, i) {
            if (i > 0) e.classList.toggle("hidden");
            e.classList.add("more");
        }));
    }));
    const shareLink = document.querySelector(".share-link");
    if (shareLink) {
        const modalTemplates = {
            share: {
                buttonSelector: '[data-modal-button="share"]',
                modalSelector: "#share-modal"
            }
        };
        Object.values(modalTemplates).forEach((modalTemplate => {
            const {buttonSelector, modalSelector} = modalTemplate;
            const modalButtons = document.querySelectorAll(buttonSelector);
            modalButtons.forEach((item => {
                item.addEventListener("click", (() => {
                    const modal = document.querySelector(modalSelector);
                    modal.classList.remove("hidden");
                    let p = modal.querySelector(".popup__content");
                    p.addEventListener("click", (e => {
                        e.stopPropagation();
                    }));
                }));
            }));
        }));
        const allModals = document.querySelectorAll("[data-modal]");
        allModals.forEach((item => {
            item.addEventListener("click", (function() {
                this.classList.add("hidden");
            }));
        }));
        const copyButton = document.getElementById("copy");
        const linkText = document.getElementById("text");
        const copyLink = document.querySelector(".copy-link");
        copyButton.addEventListener("click", (() => {
            const textToCopy = linkText.textContent.trim();
            console.log("Text copied to clipboard");
            if (navigator.clipboard) navigator.clipboard.writeText(textToCopy).then((() => {
                copyLink.textContent = "Link copied to clipboard";
            })).catch((error => {
                console.error("Could not copy text: ", error);
            })); else {
                const textarea = document.createElement("textarea");
                textarea.value = textToCopy;
                document.body.appendChild(textarea);
                textarea.select();
                document.execCommand("copy");
                document.body.removeChild(textarea);
                copyLink.textContent = "Link copied to clipboard";
            }
        }));
    }
    const btn1 = document.querySelector('input[value="1"]');
    const btn2 = document.querySelector('input[value="2"]');
    const btn3 = document.querySelector('input[value="3"]');
    const groupRadio = document.querySelector(".group-radio");
    if (groupRadio) {
        const contentAlbums = document.getElementById("content-albums");
        const contentSingle = document.getElementById("content-single");
        const contentCompilations = document.getElementById("content-compilations");
        const items1 = contentAlbums.children;
        const items2 = contentSingle.children;
        const items3 = contentCompilations.children;
        let count1 = items1.length;
        let count2 = items2.length;
        let count3 = items3.length;
        document.getElementById("count1").innerText = count1;
        document.getElementById("count2").innerText = count2;
        document.getElementById("count3").innerText = count3;
        contentAlbums.style.display = "block";
        contentSingle.style.display = "none";
        contentCompilations.style.display = "none";
        count1 -= contentAlbums.querySelectorAll(".class1").length;
        document.getElementById("count1").innerText = count1;
        document.getElementsByName("filter").forEach((btn => {
            btn.addEventListener("change", (() => {
                if (btn === btn1) {
                    contentAlbums.style.display = "block";
                    contentSingle.style.display = "none";
                    contentCompilations.style.display = "none";
                    count1 = items1.length;
                    count2 = items2.length;
                    count3 = items3.length;
                    count1 -= contentAlbums.querySelectorAll(".class1").length;
                    document.getElementById("count1").innerText = count1;
                } else if (btn === btn2) {
                    contentAlbums.style.display = "none";
                    contentSingle.style.display = "block";
                    contentCompilations.style.display = "none";
                    count1 = items1.length;
                    count2 = items2.length;
                    count3 = items3.length;
                    count2 -= contentSingle.querySelectorAll(".class1").length;
                    document.getElementById("count2").innerText = count2;
                } else if (btn === btn3) {
                    contentAlbums.style.display = "none";
                    contentSingle.style.display = "none";
                    contentCompilations.style.display = "block";
                    count1 = items1.length;
                    count2 = items2.length;
                    count3 = items3.length;
                    count3 -= contentCompilations.querySelectorAll(".class1").length;
                    document.getElementById("count3").innerText = count3;
                }
            }));
        }));
    }
    const filterGuest = document.querySelector(".filter-guest");
    if (filterGuest) {
        const filterBtn = document.querySelector(".filter-guest__btn");
        const filterList = filterGuest.querySelector(".filter-guest__list");
        filterBtn.addEventListener("click", (() => {
            filterList.classList.toggle("filter-active");
            filterBtn.classList.toggle("arrow");
        }));
        function sortList(order) {
            const lists = document.querySelectorAll(".sort__list");
            lists.forEach((list => {
                const items = list.querySelectorAll("li");
                const sortedItems = Array.from(items).sort((function(a, b) {
                    const aTitle = a.querySelector(".guest__title").textContent;
                    const bTitle = b.querySelector(".guest__title").textContent;
                    if ("asc" === order) return aTitle.localeCompare(bTitle); else return bTitle.localeCompare(aTitle);
                }));
                sortedItems.forEach((function(item) {
                    list.appendChild(item);
                }));
            }));
        }
        function sortYears(order) {
            const lists = document.querySelectorAll(".sort__list");
            lists.forEach((list => {
                const liElements = Array.from(list.children);
                liElements.sort((function(a, b) {
                    const yearA = parseInt(a.querySelector(".guest__year").textContent);
                    const yearB = parseInt(b.querySelector(".guest__year").textContent);
                    if ("asc" === order) return yearA - yearB; else return yearB - yearA;
                }));
                liElements.forEach((li => {
                    list.appendChild(li);
                }));
            }));
        }
        document.getElementById("sortAscBtn").addEventListener("click", (function() {
            sortYears("asc");
        }));
        document.getElementById("sortDescBtn").addEventListener("click", (function() {
            sortYears("desc");
        }));
        document.getElementById("sortA").addEventListener("click", (function() {
            sortList("asc");
        }));
        document.getElementById("sortZ").addEventListener("click", (function() {
            sortList("desc");
        }));
    }
})();