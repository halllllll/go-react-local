import { FC, Suspense } from "react";
import { useGetCounts } from "./service/count";
import { ErrorBoundary } from "react-error-boundary";
import { ErrorFallback } from "./error-boundary";

const CountTable: FC = () => {
  const { data: counts } = useGetCounts();
  return (
    <div>
      <table>
        <thead>
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Value</th>
            <th scope="col">created</th>
            <th scope="col">updated</th>
          </tr>
        </thead>
        <tbody>
          {counts.map((count) => {
            return (
              <tr key={count.id}>
                <td>{count.id}</td>
                <td>{count.value}</td>
                <td>{count.created}</td>
                <td>{count.updated}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};

export const CountListComponent = () => {
  return (
    <div>
      <h3>count list</h3>
      <ErrorBoundary
        FallbackComponent={ErrorFallback}
        onReset={() => {
          console.log("reset");
        }}
      >
        <Suspense fallback={<h2>wait...</h2>}>
          <CountTable />
        </Suspense>
      </ErrorBoundary>
    </div>
  );
};
