(() => {
    "use strict";
    const dotsFunc = () => {
        const dots = document.querySelector(".dots");
        if (dots) {
            const links = document.querySelectorAll(".album__link");
            if (links.length > 1) dots.style.display = "block";
            links.forEach((function (e, i) {
                if (i > 0) e.classList.add("hidden");
            }));
            dots.addEventListener("click", (function () {
                links.forEach((function (e, i) {
                    if (i > 0) e.classList.toggle("hidden");
                    e.classList.add("more");
                }));
            }));
        }
    };
    dotsFunc();
    const headerScrollFunc = () => {
        const scrollContent = document.querySelector(".scroll-content");
        if (scrollContent) {
            const header = document.querySelector("header");
            const initialHeaderPosition = header.offsetTop;
            scrollContent.addEventListener("scroll", (function () {
                const scrolledPixels = scrollContent.scrollTop;
                header.style.top = initialHeaderPosition - scrolledPixels + "px";
            }));
        }
    };
    headerScrollFunc();
    const shareLinkFunc = () => {
        const shareLink = document.querySelector(".share-link");
        if (shareLink) {
            const modalTemplates = {
                share: {
                    buttonSelector: '[data-modal-button="share"]',
                    modalSelector: "#share-modal"
                }
            };
            Object.values(modalTemplates).forEach((modalTemplate => {
                const { buttonSelector, modalSelector } = modalTemplate;
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
                item.addEventListener("click", (function () {
                    this.classList.add("hidden");
                }));
            }));
        }
    };
    shareLinkFunc();
    const copyFunc = () => {
        const copyButton = document.getElementById("copy");
        if (copyButton) {
            const linkTextElement = document.getElementById("text");
            const copyLinkElement = document.querySelector(".copy-link");
            copyButton.addEventListener("click", (() => {
                const textToCopy = linkTextElement.textContent.trim();
                const textarea = document.createElement("textarea");
                textarea.style.cssText = "position: fixed; top: -9999px;";
                textarea.value = textToCopy;
                document.body.appendChild(textarea);
                textarea.select();
                document.execCommand("copy");
                document.body.removeChild(textarea);
                copyLinkElement.textContent = "Link copied to clipboard";
            }));
        }
    };
    copyFunc();
    const filterGuestFunc = () => {
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
                    const sortedItems = Array.from(items).sort((function (a, b) {
                        const aTitle = a.querySelector(".guest__title").textContent;
                        const bTitle = b.querySelector(".guest__title").textContent;
                        if ("asc" === order) return aTitle.localeCompare(bTitle); else return bTitle.localeCompare(aTitle);
                    }));
                    sortedItems.forEach((function (item) {
                        list.appendChild(item);
                    }));
                }));
            }
            function sortYears(order) {
                const lists = document.querySelectorAll(".sort__list");
                lists.forEach((list => {
                    const liElements = Array.from(list.children);
                    liElements.sort((function (a, b) {
                        const yearA = parseInt(a.querySelector(".guest__year").textContent);
                        const yearB = parseInt(b.querySelector(".guest__year").textContent);
                        if ("asc" === order) return yearA - yearB; else return yearB - yearA;
                    }));
                    liElements.forEach((li => {
                        list.appendChild(li);
                    }));
                }));
            }
            document.getElementById("sortAscBtn").addEventListener("click", (function () {
                sortYears("asc");
            }));
            document.getElementById("sortDescBtn").addEventListener("click", (function () {
                sortYears("desc");
            }));
            document.getElementById("sortA").addEventListener("click", (function () {
                sortList("asc");
            }));
            document.getElementById("sortZ").addEventListener("click", (function () {
                sortList("desc");
            }));
        }
    };
    filterGuestFunc();
    const countFunc = () => {
        const itemsStarus = document.querySelectorAll(".guest__status");
        if (itemsStarus) {
            const counts = {};
            itemsStarus.forEach((item => {
                const content = item.textContent;
                if (counts.hasOwnProperty(content)) counts[content]++; else counts[content] = 1;
            }));
            document.querySelector("#count1").textContent = counts["album"] || 0;
            document.querySelector("#count2").textContent = counts["single"] || 0;
            document.querySelector("#count3").textContent = counts["compilation"] || 0;
        }
    };
    countFunc();
    const statusFunc = () => {
        const groupRadio = document.querySelector(".group-radio");
        if (groupRadio) {
            function filterByStatus(status) {
                const items = document.querySelectorAll(".guest__item");
                items.forEach((item => {
                    const statusDiv = item.querySelector(".guest__status");
                    if ("all" === status || statusDiv.textContent !== status) item.classList.add("hidden"); else item.classList.remove("hidden");
                }));
            }
            const radioButtons = document.querySelectorAll('input[name="filter"]');
            radioButtons.forEach((button => {
                button.addEventListener("click", (() => {
                    filterByStatus(button.value);
                }));
            }));
        }
    };
    statusFunc();
})();
