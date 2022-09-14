import axios from "axios";
import { Gift } from "../types";

const baseUrl = "http://localhost:8080"

export const getGift = async (id: string) => {
    const response = await axios.get<Gift>(`${baseUrl}/gifts/${id}`);
    return response.data;
}

export const postUseGift = async (gift: Gift) => {
    if (gift.amount <= 0) {
        return;
    }

    const response = await axios.post<Gift>(`${baseUrl}/gifts/${gift.giftId}/use`);
    return response.data;
}
