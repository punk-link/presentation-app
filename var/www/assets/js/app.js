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
    const sortListFunc = () => {
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
                filterList.classList.toggle("filter-active");
            }));
            document.getElementById("sortDescBtn").addEventListener("click", (function () {
                sortYears("desc");
                filterList.classList.toggle("filter-active");
            }));
            document.getElementById("sortA").addEventListener("click", (function () {
                sortList("asc");
                filterList.classList.toggle("filter-active");
            }));
            document.getElementById("sortZ").addEventListener("click", (function () {
                sortList("desc");
                filterList.classList.toggle("filter-active");
            }));
        }
    };
    sortListFunc();
    const filterCheckboxFunc = () => {
        const filterCheckboxes = document.getElementsByName("filter");
        if (filterCheckboxes) {
            function handleFilterChange() {
                const selectedFilters = Array.from(filterCheckboxes).filter((checkbox => checkbox.checked)).map((checkbox => checkbox.value));
                const items = document.getElementsByClassName("guest__item");
                for (let i = 0; i < items.length; i++) {
                    const item = items[i];
                    const guestStatus = item.querySelector(".guest__status").textContent;
                    if (0 === selectedFilters.length || selectedFilters.includes(guestStatus)) item.style.display = "flex"; else item.style.display = "none";
                }
            }
            for (let i = 0; i < filterCheckboxes.length; i++) filterCheckboxes[i].addEventListener("change", handleFilterChange);
        }
    };
    filterCheckboxFunc();
})();