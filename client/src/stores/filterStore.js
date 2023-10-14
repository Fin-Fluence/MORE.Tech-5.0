import { defineStore } from 'pinia';
import { ref} from 'vue';

export const useFilterStore = defineStore('filter', () => {
  const filter = ref({
    // limited: true,
    // weekend: true,
    // rent_boxes: true,
    // biometric_registration: true,
    // sims: true,
    // registration_ip: true,
    service_names: []
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

  function addServiceName(serverName) {
    filter.value.service_names.push(serverName);
  }

  function removeServiceName(serverName) {
    const index = filter.value.service_names.indexOf(serverName);
    if (index !== -1) {
      filter.value.service_names.splice(index, 1);
    }
  }

  return { filter, setFilterProperty, resetFilter, deleteFilterProperty, addServiceName, removeServiceName };
});
