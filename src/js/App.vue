<template>
    <div class="container-fluid min-vh-100 hokusai">
        <div class="row h-100 align-items-center">
            <div class="col">
                <Search />
            </div>
        </div>
        <div class="row h-100 align-items-center">
            <div class="col">
                <Actions />
            </div>
        </div>
    </div>
</template>

<script>
import Search from "./components/Search.vue";
import Actions from "./components/Actions.vue";

export default {
    components: {
        Actions,
        Search
    },
};
</script>