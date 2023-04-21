function updateSharingURL() {
    var href = window.location.href;

    var facebookLink = document.getElementsByClassName("popup__media-item facebook")[0];
    facebookLink.href = href;

    var instagramLink = document.getElementsByClassName("popup__media-item instagram")[0];
    instagramLink.href = href;

    var twitterLink = document.getElementsByClassName("popup__media-item twitter")[0];
    twitterLink.href = href;

    var vkLink = document.getElementsByClassName("popup__media-item vk")[0];
    vkLink.href = href;

    var directLink = document.getElementById("text");
    directLink.innerHTML = href + directLink.innerHTML;
}
