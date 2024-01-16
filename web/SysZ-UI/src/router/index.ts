import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'login',
            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import('@/views/LoginView.vue')
        },
        {
            path: '/workspace',
            name: 'workspace',
            component: () => import('@/views/WorkSpace.vue'),
            children: [
                {
                    path: '',
                    name: 'home',
                    component: () => import('@/components/workspace/WorkspaceHome.vue')
                },
                {
                    path: 'system',
                    name: 'system',
                    component: () => import('@/components/workspace/WorkspaceSystemStatus.vue'),
                },
                {
                    path: 'namespace',
                    name: 'namespace',
                    component: () => import('@/components/workspace/WorkSpaceNamespace.vue')
                },
                {
                    path: 'server',
                    name: 'server',
                    component: () => import('@/components/workspace/WorkspaceServer.vue')
                },
                {
                    path: 'blacklist',
                    name: 'blacklist',
                    component: () => import('@/components/workspace/WorkspaceBlacklist.vue')
                },
                {
                    path: 'operator',
                    name: 'operator',
                    component: () => import('@/components/workspace/WorkspaceOperator.vue')
                },
                {
                    path: 'document',
                    name: 'document',
                    component: () => import('@/components/workspace/WorkspaceDocument.vue')
                },
            ]
        }
    ]
})

export default router