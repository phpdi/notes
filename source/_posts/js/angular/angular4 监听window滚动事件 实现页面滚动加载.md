###错误做法
使用render2的listen方法进行监听，之前这样做，问题是，监听事件触发后，在其他组件中触发window的滚动
###正确做法
```

    //监听滚动，加载数据
    @HostListener('window:scroll', ['$event']) public onScroll = ($event) => {

        //客户端高度
        var clientH = document.documentElement.clientHeight;
        //body高度
        var bodyH = document.body.clientHeight;

        //滚动的高度
        var scrollTop = document.documentElement.scrollTop;
        console.log(bodyH)
        //滚动到底部60以内
        if (bodyH - clientH - scrollTop 

        <input type="hidden" name="content" value="###错误做法
使用render2的listen方法进行监听，之前这样做，问题是，监听事件触发后，在其他组件中触发window的滚动
###正确做法
```

    //监听滚动，加载数据
    @HostListener('window:scroll', ['$event']) public onScroll = ($event) => {

        //客户端高度
        var clientH = document.documentElement.clientHeight;
        //body高度
        var bodyH = document.body.clientHeight;

        //滚动的高度
        var scrollTop = document.documentElement.scrollTop;
        console.log(bodyH)
        //滚动到底部60以内
        if (bodyH - clientH - scrollTop < 80) {
            if (!this.flag) {
                console.log('翻页');
                //翻页
                this.changePage('+');
            }
            this.flag = true;
        } else {
            this.flag = false;
        }
    }

```
设置flag的目的是避免在滚动的过程中重复加载数据，到达只加载一次的目的">