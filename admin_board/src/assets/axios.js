import router from '@/router'
import axios from 'axios';
let axiosInstance = {
	...axios.create({
		timeout: 5000,
		withCredentials: false,
		showAlertOnFail: true,
	}),
	all: axios.all,
	spread: axios.spread,
};

axiosInstance.interceptors.request.use(
	config => {
		config.headers['Authorization'] = `Bearer ${localStorage.getItem('jwt-token')}`
		return config
	},
	err => {
		return Promise.reject(err)
	},
)

axiosInstance.interceptors.response.use(
	response => {
		return response
	},
	error => {
		if (error.response) {
			switch (error.response.status) {
				case 401:
	            // Unauthorized
	            	router.push({ path: '/' });
      		}
  		}
  		let code = '?';
        let message = '';
        console.log("DEBUG",error.response)
    	return Promise.reject(error)
    },
)

const install = (Vue, opts = {}) => {
	if (install.installed) return;
	install.installed = true;
	Vue.axios = axiosInstance;
	Object.defineProperties(Vue.prototype, {
		$axios: {
			get() {
				return axiosInstance;
			},
		},
	});
};

export default {
	axios: axiosInstance,
	install,
}