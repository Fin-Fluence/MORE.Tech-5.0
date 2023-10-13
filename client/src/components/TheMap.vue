<script setup>
import * as ymaps3 from 'ymaps3';
import markImage from '@/assets/images/icons/mark.svg';
import { ref } from 'vue';

console.log()
initMap();

async function initMap() {
    await ymaps3.ready;

    const {YMap, YMapDefaultSchemeLayer, YMapDefaultFeaturesLayer, YMapMarker} = ymaps3;

    // Координаты центра карты
    const CENTER_COORDINATES = [37.623082, 55.752540];
    const LOCATION = {center: CENTER_COORDINATES, zoom: 9};

    const map = new YMap(document.getElementById('map'), {location: LOCATION});

    const content = document.createElement('div');
    content.innerHTML = '<div>Тут может быть все что угодно</div>';
        
    // Добавляем слой для отображения схематической карты
    map.addChild(new YMapDefaultSchemeLayer())
    map.addChild(new YMapDefaultFeaturesLayer())

    const marks = ref([
        {
            coordinates: [37.623082, 55.752540],
            desc: 'test test test test test test test test',
        },
        {
            coordinates: [38.623082, 55.752540],
            desc: 'test test test test test test test test',
        },
        {
            coordinates: [39.623082, 55.752540],
            desc: 'test test test test test test test test',
        },
        {
            coordinates: [40.623082, 55.752540],
            desc: 'test test test test test test test test',
        },
        {
            coordinates: [41.623082, 55.752540],
            desc: 'test test test test test test test test',
        },
    ])
    marks.value.forEach(mark => {
        // Создание маркера
        const el = document.createElement('img');
        el.className = 'marker';
        el.src = markImage;
        el.title = 'Маркер';

        el.onclick = () => {console.log('test')}

        // Создание заголовка маркера
        const markerTitle = document.createElement('div');
        markerTitle.className = 'marker-title';

        // Контейнер для элементов маркера
        const imgContainer = document.createElement('div');
        imgContainer.appendChild(el);
        imgContainer.appendChild(markerTitle);
        const MARKER_COORDINATES = mark.coordinates
        
        // Добавление маркера на карту
        map.addChild(new YMapMarker({coordinates: MARKER_COORDINATES}, imgContainer));
    });




}

</script>

<template>
  <div id="map"></div>
</template>


<style lang="scss">
#map {
    width: 100%;
    height: 100vh;
}
</style>