import React from "react";
import {
    Card,
    CardContent,
    CardHeader,
    createStyles,
    makeStyles,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableRow,
    Theme
} from "@material-ui/core";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        card: {
            marginLeft: theme.spacing(1),
            marginRight: theme.spacing(1),
        },
        table: {
            minWidth: 650,
            marginBottom: 9
        },
    }),
);

interface Task {
    id: string,
    title: string,
    description: string,
    userName: string,
}

const rows: Array<Task> = [
    {id: 'id', title: 'title', description: 'desc', userName: 'user'},
    {id: 'id', title: 'title', description: 'desc', userName: 'user'},
    {id: 'id', title: 'title', description: 'desc', userName: 'user'},
    {id: 'id', title: 'title', description: 'desc', userName: 'user'},
];


const TasksPage: React.FC = () => {
    const classes = useStyles();
    return (
        <Card className={classes.card}>
            <CardHeader
                title={'Tasks'}
            />
            <CardContent>
                <Table className={classes.table}>
                    <TableHead>
                        <TableRow>
                            <TableCell>ID</TableCell>
                            <TableCell>title</TableCell>
                            <TableCell>description</TableCell>
                            <TableCell>user name</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rows.map(row => (
                            <TableRow key={row.id}>
                                <TableCell>{row.id}</TableCell>
                                <TableCell>{row.title}</TableCell>
                                <TableCell>{row.description}</TableCell>
                                <TableCell>{row.userName}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </CardContent>
        </Card>
    );
};

export default TasksPage;
