import Login from './views/Login.vue'
import NotFound from './views/404.vue'
import Home from './views/Home.vue'
import Main from './views/Main.vue'
import Test from './views/test.vue'
import ControlIndex from './views/control/index.vue'
import Setting from './views/setting/setting.vue'
let routes = [
    {
        path: '/login',
        component: Login,
        name: '',
        hidden: true
    },
    {
        path: '/404',
        component: NotFound,
        name: '',
        hidden: true
    },
    {
        path: '/',
        redirect: { path: '/control' }
    },
    {
        path: '/control',
        component: ControlIndex,
        name: '控制面板',
        iconCls: 'el-icon-message'//图标样式class

    },
    {
        path: '/test',
        component: Test,
        name: '测试',
        iconCls: 'el-icon-message'//图标样式class

    },
    {
        path: '/setting',
        component: Setting,
        name: '系统设置',
        iconCls: 'el-icon-message'//图标样式class

    },
    {
        path: '*',
        hidden: true,
        redirect: { path: '/404' }
    }
];

export default routes;