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
import {useUserState} from "../components/Providers/UserProvider";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        tasks: {
            display: `grid`,
            gridTemplateRows: `2000px`,
            gridTemplateColumns: `1fr 90% 1fr`,
        },
        mainSection: {
            gridColumn: `2 / 3`
        },
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
    const {signOut} = useUserState();
    return (
        <div className={classes.tasks}>
            <div className={classes.mainSection}>
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
                                {rows.map((row, index) => (
                                    <TableRow key={'tasks-table-row-' + index}>
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
            </div>
        </div>
    );
};

export default TasksPage;
