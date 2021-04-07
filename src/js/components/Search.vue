<template>
    <div class="container search">
        <div class="row">
            <div class="col">
                <form v-on:submit.prevent="submit">
                    <div class="input-group">
                        <div class="input-group-prepend engines">
                            <ul class="list-inline">
                                <li 
                                    v-for="engine in engines"
                                    :key="engine.name"
                                    class="list-inline-item"
                                >
                                    <a href="#" :class="activeEngine(engine.name)" v-on:click.prevent="selectEngine(engine.name)">
                                        <Icon :icon="engine.icon" class="icon" />
                                    </a>
                                </li>                                
                            </ul>
                        </div>
                        <input 
                            v-model="query"
                            placeholder="At your service..."
                            type="text" 
                            class="form-control" 
                            aria-label="Query" 
                            aria-describedby="button-addon2"
                            >
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>

import { Icon, addCollection } from '@iconify/vue';
import mdiIcons from '@iconify/json/json/simple-icons.json';
import icoMoonIcons from '@iconify/json/json/icomoon-free.json';
import simpleIcons from '@iconify/json/json/mdi.json';

addCollection(mdiIcons);
addCollection(simpleIcons);
addCollection(icoMoonIcons);

const engines = {
    ddg: {
        name: 'ddg',
        engine: 'https://duckduckgo.com/?q=',
        icon: 'simple-icons:duckduckgo'
    },
    so: {
        name: 'so',
        engine: 'https://stackoverflow.com/search?q=',
        icon: 'simple-icons:stackoverflow'
    },
    gh: {
        name: 'gh',
        engine: 'https://github.com/search?q=',
        icon: 'simple-icons:github'
    },
    wk: {
        name: 'wk',
        engine: 'https://en.wikipedia.org/wiki/Special:Search?search=',
        icon: 'simple-icons:wikipedia'
    }
}

export default {
    components: {
        Icon
    },
    data () {
        return {
            engines: engines,
            engine: '',
            query: ''
        }
    },
    methods: {
        submit: function () {
            const url = engines[this.engine].engine + this.query
            window.open(url,'_self')
        },
        selectEngine: function (engine) {
            localStorage.setItem('defaultEngine', engine)
            this.engine = engine
        },
        activeEngine: function (engine) {
            return {
                active: (this.engine === engine)
            }
        }
    },
    mounted: function () {
        const defaultEngine = localStorage.getItem('defaultEngine')

        if (! defaultEngine) {
            localStorage.setItem('defaultEngine', 'ddg')
        }

        this.engine = localStorage.getItem('defaultEngine')
    }
}
</script>
