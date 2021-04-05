<template>
    <div class="container search">
        <div class="row">
            <div class="col">
                <form v-on:submit.prevent="submit">
                    <div class="input-group">
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

export default {
    data () {
        return {
            query: ''
        }
    },
    methods: {
        submit () {
            const url = 'https://duckduckgo.com/?q=' + this.query
            window.open(url,'_self')
        }
    }
}
</script>
