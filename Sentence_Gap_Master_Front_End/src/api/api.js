import axios from 'axios'

export const choose = (id) => {
    return axios.get(`/choose`, {
        params: {
            id,
        },
    })
}

export const submitQuestion = (msg) => {
    return axios.post(`/submit`, msg)
}
