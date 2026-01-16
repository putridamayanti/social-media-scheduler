import {Field, FieldGroup, FieldLabel} from "@/components/ui/field";
import {Input} from "@/components/ui/input";
import {Button} from "@/components/ui/button";
import {useFormik} from "formik";
import {Textarea} from "@/components/ui/textarea";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { CalendarIcon } from "lucide-react";
import { format } from "date-fns";
import { Calendar } from "@/components/ui/calendar";
import {convertDateTimeToLocalWithFormat, joinDateAndTime} from "@/lib/helpers";
import {useRouter} from "next/navigation";

interface PostFormType {
    data: any | null,
    onClose: () => void
}

interface PostFormRequestType {
    title: string;
    content: string;
    channel: string;
    scheduled_at: Date | undefined;
    scheduled_time: string;
    status: string;
}

const channels = [
    'facebook',
    'instagram',
    'twitter',
    'linkedin'
];

const statuses = [
    'scheduled',
    'cancel',
    'published'
];

export default function PostForm(props: PostFormType) {
    const { data, onClose } = props;
    const router = useRouter();

    const formik = useFormik({
        initialValues: {
            title: data?.title ?? 'Morning',
            content: data?.content ?? 'Good morning Twitter 👋',
            channel: data?.channel ?? 'twitter',
            scheduled_at: data?.scheduled_at ?? undefined,
            scheduled_time: data?.scheduled_at ? convertDateTimeToLocalWithFormat(data?.scheduled_at, 'hh:mm:ss') : '00:00:00',
            status: 'scheduled'
        },
        enableReinitialize: true,
        onSubmit: (values: PostFormRequestType) => handleSubmit(values)
    })

    const handleSubmit = async (values: PostFormRequestType) => {
        if (values.scheduled_at) {
            values.scheduled_at = joinDateAndTime(values.scheduled_at, values.scheduled_time);
        }

        const resp = await fetch(data?.id ? `/api/posts/${data?.id}` : '/api/posts', {
            method: data?.id ? 'PUT' : 'POST',
            body: JSON.stringify(values)
        });
        if (resp.status === 200) {
            onClose();
            router.refresh();
        }
    };

    return (
        <form onSubmit={formik.handleSubmit} className="pt-8">
            <FieldGroup>
                <Field>
                    <FieldLabel htmlFor="title">Title</FieldLabel>
                    <Input
                        id="title"
                        required
                        name="title"
                        onChange={formik.handleChange}
                        value={formik.values.title}
                    />
                </Field>
                <Field>
                    <FieldLabel htmlFor="content">Content</FieldLabel>
                    <Textarea
                        id="content"
                        required
                        rows={6}
                        name="content"
                        onChange={formik.handleChange}
                        value={formik.values.content}
                    />
                </Field>
                <Field>
                    <FieldLabel>Channel</FieldLabel>
                    <div className="space-x-3">
                        {channels.map((e, i) => (
                            <Button
                                key={i}
                                type="button"
                                className="capitalize"
                                variant={formik.values.channel === e ? 'default' : 'outline'}
                                size="sm"
                                onClick={() => formik.setFieldValue('channel', e)}>
                                {e}
                            </Button>
                        ))}
                    </div>
                </Field>
                <Field>
                    <FieldLabel htmlFor="scheduled_at">Schedule for</FieldLabel>
                    <div className="flex gap-4">
                        <Popover>
                            <PopoverTrigger asChild>
                                <Button
                                    variant="outline"
                                    type="button"
                                    data-empty={!formik.values.scheduled_at}
                                    className="data-[empty=true]:text-muted-foreground w-[280px] justify-start text-left font-normal"
                                >
                                    <CalendarIcon />
                                    {formik.values.scheduled_at ? format(formik.values.scheduled_at, "PPP") : <span>Schedule for</span>}
                                </Button>
                            </PopoverTrigger>
                            <PopoverContent className="w-auto p-0">
                                <Calendar
                                    mode="single"
                                    selected={formik.values.scheduled_at}
                                    onSelect={(date) => formik.setFieldValue('scheduled_at', date)} />
                            </PopoverContent>
                        </Popover>
                        <Input
                            type="time"
                            id="time-picker"
                            step="1"
                            defaultValue="10:30:00"
                            onChange={(e) => formik.setFieldValue('scheduled_time', e.target.value)}
                            value={formik.values.scheduled_time}
                        />
                    </div>
                    <div className="text-muted-foreground px-1 text-sm">
                        Your post will be published on{" "}
                        <span className="font-medium">{formik.values.scheduled_at ? format(formik.values.scheduled_at, "PPP") : ''}</span>&nbsp;
                        at <span>{formik.values.scheduled_time ?? ''}</span>.
                    </div>
                </Field>
                <Field>
                    <FieldLabel>Channel</FieldLabel>
                    <div className="space-x-3">
                        {statuses.map((e, i) => (
                            <Button
                                key={i}
                                type="button"
                                className="capitalize"
                                variant={formik.values.status === e ? 'default' : 'outline'}
                                size="sm"
                                onClick={() => formik.setFieldValue('status', e)}>
                                {e}
                            </Button>
                        ))}
                    </div>
                </Field>
                <Field>
                    <Button type="submit">Submit</Button>
                </Field>
            </FieldGroup>
        </form>
    )
}
