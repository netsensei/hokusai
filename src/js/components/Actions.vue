<template>
    <div class="container actions">
        <div class="row justify-content-md-center">
            <div 
                v-for="actionGroup in actionGroups"
                :key="actionGroup.name"
                class="action-group col-lg-4"
            >
                <h2 class="action-group-title">{{ actionGroup.label }}</h2>
                <div class="action-grid">
                <a
                    v-for="action in actionGroup.actions"
                    :key="action.name"
                    class="action"
                    :href="action.link"
                >
                    <div class="action-icon">
                        <Icon :icon="action.icon" class="icon" />
                    </div>
                </a>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

import { Icon, addCollection } from '@iconify/vue';
import mdiIcons from '@iconify/json/json/simple-icons.json';
import icoMoonIcons from '@iconify/json/json/icomoon-free.json';
import simpleIcons from '@iconify/json/json/mdi.json';


const actionGroups = [
    {
        name: "comms",
        label: "Comms",
        actions: [
            {
                icon: "mdi:youtube",
                name: "YouTube",
                link: "https://youtube.com"
            },
            {
                icon: "mdi:reddit",
                name: "Reddit",
                link: "https://reddit.com"
            },
            {
                icon: "icomoon-free:hackernews",
                name: "Hackernews",
                link: "https://news.ycombinator.com/"
            },
            {
                icon: "simple-icons:mastodon",
                name: "Mastodon",
                link: "https://mastodon.social"
            },
        ]
    },
    {
        name: "coding",
        label: "Coding",
        actions: [
            {
                icon: "simple-icons:github",
                name: "Github",
                link: "https://github.com"
            },
            {
                icon: "simple-icons:bitbucket",
                name: "Bitbucket",
                link: "https://bitbucket.com"
            },
        ]
    },
    {
        name: "productivity",
        label: "Productivity",
        actions: [
            {
                icon: "simple-icons:whatsapp",
                name: "Whatsapp",
                link: "https://web.whatsapp.com/"
            },
            {
                icon: "simple-icons:gmail",
                name: "Gmail",
                link: "https://gmail.com"
            },
            {
                icon: "simple-icons:googledrive",
                name: "Gmail",
                link: "https://drive.google.com/"
            },
        ]
    },
]

addCollection(mdiIcons);
addCollection(simpleIcons);
addCollection(icoMoonIcons);

export default {
    components: {
        Icon
    },
    data () {
        return {
            actionGroups: actionGroups
        }
    }
}

</script>


