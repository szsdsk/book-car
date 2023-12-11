const Left = -3, Right = 3;
let itemsPerPage = 5;
let currentPage = 1;
let data = document.querySelectorAll('.probation-')
let curTable = 'probation'
let totalPages = Math.ceil(data.length / itemsPerPage);
let size = data.length;
let map = new Map();
map.set('probation', '#itemsPerPage1')
map.set('location', '#itemsPerPage2')
map.set('trend', '#itemsPerPage3')
map.set('probationContainer', '#pagination-container1')
map.set('locationContainer', '#pagination-container2')
map.set('trendContainer', '#pagination-container3')

// 事件委托，判断点击事件，转到相应的表
document.querySelector('.title').addEventListener('click', event => {
    itemsPerPage = 5;
    currentPage = 1;
    let li = event.target.closest('li');
    if (!li)
        return;
    if (!document.querySelector('.title').contains(li))
        return;
    curTable = li.dataset.table;
    data = document.querySelectorAll(li.dataset.class);
    size = data.length;
    totalPages = Math.ceil(size / itemsPerPage)
    init();
});

// 显示某一个表
function activeTable() {
    const titles = document.querySelector('.title').children;
    for (const title of titles) {
        let name = title.querySelector('span').textContent.toLowerCase();
        if (name.includes(curTable.toLowerCase())) {
            title.classList.add('active');
        } else {
            title.classList.remove('active');
        }
    }
}


// 初始化页面
function init() {
    for (let child of document.querySelector('.content').children) {
        if (child.className !== curTable) {
            child.style.display = 'none';
        } else {
            child.style.display = 'block';
        }
    }
    activeTable();
    document.querySelector(map.get(curTable)).firstElementChild.selected  = 'true';
    showPage(currentPage); // 默认显示第一页
    renderPagination(); // 渲染分页按钮
}

// 隐藏所有表。
function hide() {
    data.forEach(li => {
        li.style.display = 'none';
    })
}
// 显示表的相应数据
function showPage(pageNumber) {
    currentPage = pageNumber;
    console.log(data.length);
    const startIndex = (pageNumber - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    hide();
    for (let i = startIndex; i <= endIndex; ++i) {
        if (data[i]) {
            data[i].style.display = 'flex';
        }
    }
}

// 根据数据条数渲染页码
function renderPagination() {
    const paginationContainer = document.querySelector(map.get(curTable+'Container'));
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

// 增加链接
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
    document.querySelector(map.get(curTable+'Container')).append(listItem);
}

// 页数太多用...表示
function addEllipsis(option) {
    const ellipsis = document.createElement('a');
    ellipsis.textContent = '...';
    ellipsis.href = '#';
    ellipsis.addEventListener('click', function (event) {
        event.preventDefault();
        showPage(currentPage + option);
        renderPagination();
    });
    document.querySelector(map.get(curTable+'Container')).appendChild(document.createElement('li')).appendChild(ellipsis);
}

// 更新活动页按钮样式
function updateActivePage(activePage) {
    const paginationContainer = document.querySelector(map.get(curTable+'Container'));
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
    itemsPerPage = parseInt(document.querySelector(map.get(curTable)).value);
    totalPages = Math.ceil(size / itemsPerPage);
    showPage(1); // 切换每页显示数量后，回到第一页
    renderPagination(); // 重新渲染分页按钮
}

// 筛选是否是学生
function isStudent() {
    hide();
    let student = document.querySelector('.student').value;
    let rawData = document.querySelectorAll('.trends');
    if (student === 'Student?') {
        data = rawData;
    }
    else {
        data = Array.from(rawData).filter(st => st.querySelector('.col-3').textContent === student);
    }
    size = data.length;
    totalPages = Math.ceil(size / itemsPerPage);
    showPage(1);
    renderPagination();
}

init();

// 更新价格就是发送一个请求。
document.querySelector('#increase').addEventListener('click', async ev => {
    await fetch('/admin/increase', {method:'PATCH'});
})