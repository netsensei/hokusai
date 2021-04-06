<template>
    <div class="container search">
        <div class="row">
            <div class="col">
                <form v-on:submit.prevent="submit">
                    <div class="input-group">
                        <div class="input-group-prepend engines">
                            <ul class="list-inline">
                                <li class="list-inline-item">
                                    <a href="#" :class="classGoogle" v-on:click.prevent="selectEngine('google')">
                                        <Icon icon="mdi:google" class="icon" />
                                    </a>
                                </li>
                                <li class="list-inline-item">
                                    <a href="#" :class="classDdg" v-on:click.prevent="selectEngine('ddg')">
                                        <Icon icon="simple-icons:duckduckgo" class="icon" />
                                    </a>
                                </li>
                                <li class="list-inline-item">
                                    <a href="#" :class="classSo" v-on:click.prevent="selectEngine('so')">
                                        <Icon icon="simple-icons:stackoverflow" class="icon" />
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
                        <div class="input-group-append">
                            <button 
                            class="btn btn-outline-secondary"
                            type="button"
                            id="button-addon2"
                            v-on:click="submit">
                            Go!
                            </button>
                        </div>
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

export default {
    components: {
        Icon
    },
    data () {
        return {
            engine: 'ddg',
            query: ''
        }
    },
    computed: {
        classDdg: function () {
            return {
                active: (this.engine === 'ddg')
            }
        },
        classGoogle: function () {
            return {
                active: (this.engine === 'google')
            }
        },
        classSo: function () {
            return {
                active: (this.engine === 'so')
            }
        },
    },
    methods: {
        submit: function () {
            const engines = {
                ddg: 'https://duckduckgo.com/?q=',
                google: 'https://www.google.com/search?q=',
                so: 'https://stackoverflow.com/search?q='
            }
            
            const url = engines[this.engine] + this.query
            window.open(url,'_self')
        },
        selectEngine: function (engine) {
            this.engine = engine
        }
    }
}
</script>
