interface ISomething {
    id: string;
    priority: number;
}
class Queue {
    current_queue: Array<ISomething> = [];

    add<T extends ISomething>(something: T): void {
        this.current_queue.push(something);
    }

    process() {
        const priority_num = [];
        const priority = {};

        for (let i = 0; i < this.current_queue.length; i++) {
            for (const [index, item] of this.current_queue.entries()) {
                let temp = item;
                if (index < this.current_queue.length - 1) {
                    if (
                        this.current_queue[index + 1].priority < temp.priority
                    ) {
                        this.current_queue[index] =
                            this.current_queue[index + 1];
                        this.current_queue[index + 1] = temp;
                    }
                }
            }
        }

        for (let i = 0; i < this.current_queue.length; i++) {
            if (i !== 0) {
                if (
                    this.current_queue[i].priority !==
                    this.current_queue[i - 1].priority
                ) {
                    priority_num.push(this.current_queue[i].priority);
                    (priority[
                        this.current_queue[i].priority as keyof typeof priority
                    ] as Array<ISomething>) = [
                        ...(priority[
                            this.current_queue[i]
                                .priority as keyof typeof priority
                        ] || []),
                        this.current_queue[i]
                    ];
                } else if (
                    this.current_queue[i].priority ===
                    this.current_queue[i - 1].priority
                ) {
                    (
                        priority[
                            priority_num[
                                priority_num.length - 1
                            ] as keyof typeof priority
                        ] as Array<ISomething>
                    ).push(this.current_queue[i]);
                }
            } else {
                priority_num.push(this.current_queue[i].priority);
                (priority[
                    this.current_queue[i].priority as keyof typeof priority
                ] as Array<ISomething>) = [this.current_queue[i]];
            }
        }

        for (const p in priority) {
            for (
                let i = 0;
                i <
                (priority[p as keyof typeof priority] as Array<ISomething>)
                    .length;
                i++
            ) {
                for (const [index, item] of (
                    priority[p as keyof typeof priority] as Array<ISomething>
                ).entries()) {
                    const temp = item;

                    if (
                        index <
                        (
                            priority[
                                p as keyof typeof priority
                            ] as Array<ISomething>
                        ).length -
                            1
                    ) {
                        if (
                            ((
                                priority[
                                    p as keyof typeof priority
                                ] as Array<ISomething>
                            )[index + 1].id as unknown as string) < item.id
                        ) {
                            (
                                priority[
                                    p as keyof typeof priority
                                ] as Array<ISomething>
                            )[index] =
                                priority[p as keyof typeof priority][index + 1];
                            (
                                priority[
                                    p as keyof typeof priority
                                ] as Array<ISomething>
                            )[index + 1] = temp;
                        }
                    }
                }
            }
        }

        this.current_queue = [];

        for (const p in priority) {
            this.current_queue = [
                ...this.current_queue,
                ...priority[p as keyof typeof priority]
            ];
        }

        const temp = this.current_queue;
        this.current_queue = [];
        for (const [i, t] of temp.entries()) {
            if (i !== 0) this.current_queue.push(t);
        }
        console.log(temp[0]);
        return temp[0];
    }

    print() {
        console.log(this.current_queue);
    }
}

const queue = new Queue();
queue.add<ISomething>({ id: "T3", priority: 2 });
queue.add<ISomething>({ id: "T1", priority: 1 });
queue.add<ISomething>({ id: "T6", priority: 3 });
queue.add<ISomething>({ id: "T5", priority: 1 });
queue.add<ISomething>({ id: "T4", priority: 2 });
queue.add<ISomething>({ id: "T2", priority: 1 });
queue.add<ISomething>({ id: "T7", priority: 4 });
queue.print();
queue.process();
queue.print();
queue.process();
queue.print();
