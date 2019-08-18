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
    Theme,
    Typography
} from "@material-ui/core";
import {useGetTaskQuery} from "../generated/graphql";
import {useUserState} from "../components/Providers/UserProvider";
import {Redirect} from "react-router";
import Routes from "../app/routes";

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


const TasksPage: React.FC = () => {
    const classes = useStyles();
    const {isLoggedIn} = useUserState();
    const {data, error, loading} = useGetTaskQuery({fetchPolicy: "cache-and-network"});
    if (loading) return <Typography>Loading...</Typography>;
    if (error) return <Typography>Error</Typography>;
    const tasks = data!.tasks.edges.map((edge) => {
        const node = edge.node;
        const task: Task = {
            id: node.id,
            title: node.title,
            description: node.description,
            userName: node.user ? node.user.name : ''
        };
        return task
    });
    return (
        <div className={classes.tasks}>
            <div className={classes.mainSection}>
                {isLoggedIn ? (
                    <CardTable tasks={tasks}/>
                ) : (
                    <Redirect to={Routes.signIn()}/>
                )}
            </div>
        </div>
    );
};

const CardTable = (props: { tasks: Array<Task>, }) => {
    const {tasks} = props;
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
                        {tasks.map((row, index) => (
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
    )
}

export default TasksPage;
