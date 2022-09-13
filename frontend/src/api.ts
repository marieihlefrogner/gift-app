import axios from "axios";
import {Gift} from "./types";

const baseUrl = "http://localhost:8080"

export const getGift = async (id: number) => {
    const response = await axios.get<Gift>(`${baseUrl}/gifts/${id}`);
    return response.data;
}

export const decreaseGiftAmount = async (gift: Gift) => {
    if (gift.amount <= 0) {
        return;
    }

    const update: Gift = {
        ...gift,
        amount: gift.amount-1,
    };

    const response = await axios.put<Gift>(`${baseUrl}/gifts/${gift.ID}`, update);
    return response.data;
}
