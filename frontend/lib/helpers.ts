import {format, toZonedTime} from 'date-fns-tz';
import {set} from "date-fns";

export const convertDateTimeToLocal = (date: string) => {
    const userTimeZone = Intl.DateTimeFormat().resolvedOptions().timeZone;

    const localDate = toZonedTime(date, userTimeZone);

    return format(localDate, 'PPP HH:mm:ss', {
        timeZone: userTimeZone,
    })
}

export const convertDateTimeToLocalWithFormat = (date: string, formatDate: string) => {
    const userTimeZone = Intl.DateTimeFormat().resolvedOptions().timeZone;

    const localDate = toZonedTime(date, userTimeZone);

    return format(localDate, formatDate, {
        timeZone: userTimeZone,
    })
}

export const joinDateAndTime = (date: Date, time: string) => {
    const [hours, minutes, seconds = 0] = time.split(':').map(Number);

    return set(date, {
        hours,
        minutes,
        seconds,
        milliseconds: 0,
    });
}
