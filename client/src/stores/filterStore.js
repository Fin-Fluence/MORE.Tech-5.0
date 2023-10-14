import { defineStore } from 'pinia';
import { ref} from 'vue';

export const useFilterStore = defineStore('filter', () => {
  const filter = ref({
    limited: false,
    weekend: false,
    rent_boxes: false,
    biometric_registration: false,
    sims: false,
    registration_ip: false
  });

  function setFilterProperty(propertyName, value) {
      filter.value[propertyName] = value;
  }

  function resetFilter() {
    filter.value.limited = false
    filter.value.weekend = false
  }

  function deleteFilterProperty(propertyName) {
    if (propertyName in filter.value) {
      delete filter.value[propertyName];
    }
  }

  return { filter, setFilterProperty, resetFilter, deleteFilterProperty };
});
