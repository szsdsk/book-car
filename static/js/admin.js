const Left = -3, Right = 3;
let itemsPerPage = 5;
let currentPage = 1;
const data = document.querySelectorAll('.trends')
let totalPages = Math.ceil(data.length / itemsPerPage)
// 初始化页面
function init() {
    showPage(currentPage); // 默认显示第一页
    renderPagination(); // 渲染分页按钮
}

function hide() {
    data.forEach(li => {
        li.style.display = 'none';
    })
}
function showPage(pageNumber) {
    currentPage = pageNumber;
    const startIndex = (pageNumber - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    hide();
    for (let i = startIndex; i <= endIndex; ++i) {
        if (data[i]) {
            data[i].style.display = 'flex';
        }
    }
}

function renderPagination() {
    const paginationContainer = document.getElementById('pagination-container');
    paginationContainer.innerHTML = '';
    if (totalPages <= 7) {
        for (let i = 1; i <= totalPages; i++) {
            addPageLink(i);
        }
    } else {
        // 如果总页数大于7，显示部分页码，并用 "..." 表示省略的页码
        if (currentPage <= 3) {
            for (let i = 1; i <= 5; ++i) {
                addPageLink(i);
            }
            addEllipsis(Right);
            addPageLink(totalPages);
        } else if (currentPage >= totalPages - 2) {
            // 当前页码在最后三个，显示后五个页码
            addEllipsis(1);
            addEllipsis(Left);
            for (let i = totalPages - 4; i <= totalPages; i++) {
                addPageLink(i);
            }
        } else {
            // 当前页码在中间，显示当前页码和前后两个页码
            addPageLink(1);
            addEllipsis(Left);
            for (let i = currentPage - 1; i <= currentPage + 1; i++) {
                addPageLink(i);
            }
            addEllipsis(Right);
            addPageLink(totalPages);
        }
    }
    updateActivePage(currentPage);
}

function addPageLink(pageNumber) {
    const pageLink = document.createElement('a');
    pageLink.href = '#';
    pageLink.textContent = pageNumber;
    pageLink.addEventListener('click', function (event) {
        event.preventDefault();
        showPage(pageNumber);
        renderPagination();
    });

    const listItem = document.createElement('li');
    listItem.append(pageLink);
    document.querySelector('#pagination-container').append(listItem);
}

function addEllipsis(option) {
    const ellipsis = document.createElement('a');
    ellipsis.textContent = '...';
    ellipsis.href = '#';
    ellipsis.addEventListener('click', function (event) {
        event.preventDefault();
        showPage(currentPage + option);
        renderPagination();
    });
    document.querySelector('#pagination-container').appendChild(document.createElement('li')).appendChild(ellipsis);
}

// 更新活动页按钮样式
function updateActivePage(activePage) {
    const paginationContainer = document.querySelector('#pagination-container');
    const pageLinks = paginationContainer.querySelectorAll('a');

    for (let i = 0; i < pageLinks.length; i++) {
        const pageLink = pageLinks[i];
        const pageNumber = parseInt(pageLink.textContent);

        if (pageNumber === activePage) {
            pageLink.classList.add('active');
        } else {
            pageLink.classList.remove('active');
        }
    }
}

// 处理每页显示数量变化
function onItemsPerPageChange() {
    itemsPerPage = parseInt(document.querySelector('#itemsPerPage').value);
    totalPages = Math.ceil(data.length / itemsPerPage);
    showPage(1); // 切换每页显示数量后，回到第一页
    renderPagination(); // 重新渲染分页按钮
}
init();