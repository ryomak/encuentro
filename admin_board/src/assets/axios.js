
import axios from 'axios';
let axiosInstance = {
    ...axios.create({
        timeout: 5000,
        withCredentials: false,
        showAlertOnFail: true,
    }),
    all: axios.all,
    spread: axios.spread,
    CancelToken: axios.CancelToken,
};

axiosInstance.interceptors.request.use(
	config => {
	  config.headers['Authorization']    = `Bearer ${localStorage.getItem('jwt-token')}`
	  return config
	},
	err =>{
		if (err.response.status == 401){
			redirect("/")
		}
	}
);

export default {
    axios: axiosInstance,
};
