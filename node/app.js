/*******************Vue*******************/
import Vue from 'vue';
import Chat from "./components/template.vue"
Vue.filter("dateForm", function (time) {
    if (time) {
        let dt = new Date(time);
        let checktime = dt.getTime();
        let nowtime = new Date().getTime();
        let result = (nowtime - checktime) / 1000;
        if (result <= 60) {
            return "刚刚";
        } else if (result > 60 && result <= 3600) {
            return Math.ceil((result / (60))) + '分钟前';
        } else if (result > 3600 && result <= 86400) {
            return Math.ceil((result / (3600))) + '小时前';
        } else if (result > 86400 && result <= 2592000) {
            return Math.ceil((result / (86400))) + '天前';
        } else if (result > 2592000 && result <= 31104000) {
            return Math.ceil((result / (2592000))) + '个月前';
        } else {
            var y = dt.getFullYear();
            var m = dt.getMonth() + 1;
            var d = dt.getDate();
            return y + "年" + m + "月" + d + "日";
        }
        return result;
    } else {
        return null;
    }
})
new Vue({
    el: '#c-chat',

    data: {
    },
    created() {
    },
    components: {
        "chat": Chat
    },
});
